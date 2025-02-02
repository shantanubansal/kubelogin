package oauth2cli

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ShellySavePort(state string, port int, c *Config) error {
	userUid := ""
	if os.Getenv("OIDC_USER_UID") != "" {
		userUid = os.Getenv("OIDC_USER_UID")
	}
	shellyUrl := os.Getenv("SHELLY_URL")
	if shellyUrl == "" {
		shellyUrl = "http://localhost:8080/v1/shelly"
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url := fmt.Sprintf("%s/oidc/port/save?state=%v&port=%v&userUid=%v",
		shellyUrl, state, port, userUid)
	c.Logf(fmt.Sprintf("url connecting %v", url))

	resp, err := http.Get(url)
	if err != nil {
		c.Logf("Got Errored %v", err.Error())
		return err
	}
	if resp.StatusCode != 204 {
		errorMsg := fmt.Sprintf("unable to save the port for %v and state %v", port, state)
		if resp.Body != nil {
			read, err := io.ReadAll(resp.Body)
			if err == nil {
				errorMsg = string(read)
			}
		}

		return fmt.Errorf("%v: %v", resp.StatusCode, errorMsg)
	}
	if resp == nil {
		c.Logf("got nil response from the shelly.")
	} else {
		c.Logf("got '%s' response from shelly connect", resp.StatusCode)
	}
	return nil
}
