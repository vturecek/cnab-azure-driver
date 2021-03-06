package azure

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLoginInfo(t *testing.T) {

	testcases := []struct {
		name         string
		envVarsToGet []string
		valueToCheck LoginType
	}{
		// to execute tests set environment vars using names in envVarsToGet preceded with TEST_
		{"Get LoginInfo using Service Principal", []string{"CNAB_AZURE_CLIENT_ID", "CNAB_AZURE_CLIENT_SECRET", "CNAB_AZURE_TENANT_ID"}, ServicePrincipal},
		{"Get LoginInfo using Device Code", []string{"CNAB_AZURE_APP_ID"}, DeviceCode},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {
			var config = map[string]string{
				"CNAB_AZURE_CLIENT_ID":     "",
				"CNAB_AZURE_CLIENT_SECRET": "",
				"CNAB_AZURE_TENANT_ID":     "",
				"CNAB_AZURE_APP_ID":        "",
			}
			params := 0
			// Check for environment variables to use for login these are expected to be the name of the relevant driver variable prefixed with TEST_
			for _, e := range tc.envVarsToGet {
				envvar := os.Getenv(fmt.Sprintf("TEST_%s", e))
				if len(envvar) > 0 {
					config[e] = envvar
					params++
				}
			}
			if params == len(tc.envVarsToGet) {
				t.Log("Testing Login in with Service Principal")
				loginInfo, err := LoginToAzure(config["CNAB_AZURE_CLIENT_ID"], config["CNAB_AZURE_CLIENT_SECRET"], config["CNAB_AZURE_TENANT_ID"], config["CNAB_AZURE_APP_ID"])
				assert.NoError(t, err)
				assert.Equal(t, ServicePrincipal, loginInfo.LoginType, "Expected Login type to be %v", ServicePrincipal)
			} else {
				t.Skip("Skipping Test Not All Environment Variables Set")
			}
		})

	}
}
