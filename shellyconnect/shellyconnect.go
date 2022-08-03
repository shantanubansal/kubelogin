package shellyconnect

import (
	"fmt"
	"io"
	"net/http"
)

func ShellySavePort(state string, port int) error {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/v1/shelly/oidc/save/port?state=%v&port=%v", state, port))
	if err != nil {
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
	return nil
}
