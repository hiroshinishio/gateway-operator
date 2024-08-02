package image_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"oras.land/oras-go/v2/registry/remote/auth"
	"oras.land/oras-go/v2/registry/remote/credentials"

	"github.com/kong/gateway-operator/controller/kongplugininstallation/image"
)

func TestCredentialsStoreFromString(t *testing.T) {
	testCases := []struct {
		name                string
		credentials         string
		expectedErrorMsg    string
		expectedCredentials func(t *testing.T, cs credentials.Store)
	}{
		{
			name:             "invalid credentials",
			credentials:      "foo",
			expectedErrorMsg: "invalid config format:",
		},
		{
			name: "valid credentials",
			// Field auth is base64 encoded "test:test".
			credentials: `
			{
 			  "auths": {
 			    "ghcr.io": {
 			      "auth": "dGVzdDp0ZXN0"
 			    }
 			  }
			}`,
			expectedCredentials: func(t *testing.T, cs credentials.Store) {
				t.Helper()
				require.NotNil(t, cs)
				c, err := cs.Get(context.Background(), "ghcr.io")
				require.NoError(t, err)
				require.Equal(t, auth.Credential{Username: "test", Password: "test", RefreshToken: "", AccessToken: ""}, c)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			credsStore, err := image.CredentialsStoreFromString(tc.credentials)
			if tc.expectedCredentials != nil {
				tc.expectedCredentials(t, credsStore)
			} else {
				require.ErrorContains(t, err, tc.expectedErrorMsg)
			}
		})
	}
}

func TestFetchPluginContent(t *testing.T) {
	t.Run("invalid image URL", func(t *testing.T) {
		_, err := image.FetchPluginContent(context.Background(), "foo bar", nil)
		require.ErrorContains(t, err, "unexpected format of image url: could not parse reference: foo bar")
	})

	const registryUrl = "northamerica-northeast1-docker.pkg.dev/k8s-team-playground/plugin-example/"

	t.Run("valid image", func(t *testing.T) {
		plugin, err := image.FetchPluginContent(
			context.Background(), registryUrl+"valid", nil,
		)
		require.NoError(t, err)
		require.Equal(t, plugin, []byte("plugin-content\n"))
	})

	t.Run("invalid image - to many layers", func(t *testing.T) {
		_, err := image.FetchPluginContent(
			context.Background(), registryUrl+"invalid-layers", nil,
		)
		require.ErrorContains(t, err, "expected exactly one layer with plugin, found 2 layers")
	})

	t.Run("invalid image - to many layers", func(t *testing.T) {
		_, err := image.FetchPluginContent(
			context.Background(), registryUrl+"invalid-layers", nil,
		)
		require.ErrorContains(t, err, "expected exactly one layer with plugin, found 2 layers")
	})

	t.Run("invalid image - invalid name of plugin inside of it", func(t *testing.T) {
		_, err := image.FetchPluginContent(
			context.Background(), registryUrl+"invalid-name", nil,
		)
		require.ErrorContains(t, err, `file "plugin.lua" not found in the image`)
	})
}
