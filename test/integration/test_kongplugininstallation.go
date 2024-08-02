package integration

import (
	"strings"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kong/gateway-operator/api/v1alpha1"
	"github.com/kong/gateway-operator/test/helpers"
)

func TestKongPluginInstallationEssentials(t *testing.T) {
	t.Parallel()

	namespace, cleaner := helpers.SetupTestEnv(t, GetCtx(), GetEnv())

	const registryUrl = "northamerica-northeast1-docker.pkg.dev/k8s-team-playground/plugin-example/"
	t.Log("deploying an invalid KongPluginInstallation resource")
	const kpiName = "test-kpi"
	kpiNamespace := namespace.Name
	kpi := &v1alpha1.KongPluginInstallation{
		ObjectMeta: metav1.ObjectMeta{
			Name: kpiName,
		},
		Spec: v1alpha1.KongPluginInstallationSpec{
			Image: registryUrl + "invalid-layers",
		},
	}
	kpi, err := GetClients().OperatorClient.ApisV1alpha1().KongPluginInstallations(kpiNamespace).Create(GetCtx(), kpi, metav1.CreateOptions{})
	require.NoError(t, err)
	cleaner.Add(kpi)

	t.Log("waiting for the KongPluginInstallation resource to be rejected, because of the invalid image")
	checkKongPluginInstallationConditions(
		t,
		kpiNamespace,
		kpiName,
		metav1.ConditionFalse,
		`problem with the image: "northamerica-northeast1-docker.pkg.dev/k8s-team-playground/plugin-example/invalid-layers" error: expected exactly one layer with plugin, found 2 layers`)

	t.Log("updating KongPluginInstallation resource to a valid image")
	kpi, err = GetClients().OperatorClient.ApisV1alpha1().KongPluginInstallations(kpiNamespace).Get(GetCtx(), kpiName, metav1.GetOptions{})
	kpi.Spec.Image = registryUrl + "valid"
	require.NoError(t, err)
	_, err = GetClients().OperatorClient.ApisV1alpha1().KongPluginInstallations(kpiNamespace).Update(GetCtx(), kpi, metav1.UpdateOptions{})
	require.NoError(t, err)

	t.Log("waiting for the KongPluginInstallation resource to be accepted")
	checkKongPluginInstallationConditions(t, kpiNamespace, kpiName, metav1.ConditionTrue, "plugin successfully saved in cluster as ConfigMap")

	var respectiveCM corev1.ConfigMap
	t.Log("check creation of respective ConfigMap")
	require.EventuallyWithT(t, func(c *assert.CollectT) {
		configMaps, err := GetClients().K8sClient.CoreV1().ConfigMaps(namespace.Name).List(GetCtx(), metav1.ListOptions{})
		if !assert.NoError(c, err) {
			return
		}
		var found bool
		respectiveCM, found = lo.Find(configMaps.Items, func(cm corev1.ConfigMap) bool {
			return strings.HasPrefix(cm.Name, kpiName)
		})
		if !assert.True(c, found) {
			return
		}
	}, 15*time.Second, time.Second)

	t.Log("found respective ConfigMap:", respectiveCM.Name)
	pluginContent, ok := respectiveCM.Data[kpiName+".lua"]
	require.True(t, ok, "plugin.lua not found in ConfigMap")
	require.EqualValues(t, "plugin-content\n", pluginContent)

	t.Log("delete KongPluginInstallation resource")
	err = GetClients().OperatorClient.ApisV1alpha1().KongPluginInstallations(kpiNamespace).Delete(GetCtx(), kpiName, metav1.DeleteOptions{})
	require.NoError(t, err)

	t.Log("check deletion of respective ConfigMap")
	require.EventuallyWithT(t, func(c *assert.CollectT) {
		_, err := GetClients().K8sClient.CoreV1().ConfigMaps(kpiNamespace).Get(GetCtx(), respectiveCM.Name, metav1.GetOptions{})
		assert.True(c, apierrors.IsNotFound(err), "ConfigMap not deleted")
	}, 15*time.Second, time.Second)
}

func checkKongPluginInstallationConditions(
	t *testing.T,
	namespace string,
	name string,
	conditionStatus metav1.ConditionStatus,
	expectedMessage string,
) {
	t.Helper()

	require.EventuallyWithT(t, func(c *assert.CollectT) {
		kpi, err := GetClients().OperatorClient.ApisV1alpha1().KongPluginInstallations(namespace).Get(GetCtx(), name, metav1.GetOptions{})
		if !assert.NoError(c, err) {
			return
		}
		if !assert.NotEmpty(c, kpi.Status.Conditions) {
			return
		}
		status := kpi.Status.Conditions[0]
		assert.EqualValues(c, v1alpha1.KongPluginInstallationConditionStatusAccepted, status.Type)
		assert.EqualValues(c, conditionStatus, status.Status)
		if conditionStatus == metav1.ConditionTrue {
			assert.EqualValues(c, v1alpha1.KongPluginInstallationReasonReady, status.Reason)
		} else {
			assert.EqualValues(c, v1alpha1.KongPluginInstallationReasonFailed, status.Reason)
		}
		assert.EqualValues(c, expectedMessage, status.Message)
	}, 15*time.Second, time.Second)
}
