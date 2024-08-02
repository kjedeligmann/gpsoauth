// Ad hoc Go client package for Google Play Services OAuth.
package gpsoauth

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	authURL          = "https://android.clients.google.com/auth"
	userAgent        = "GoogleAuth/1.4"
	DefaultClientSig = "38918a453d07199354f8b19af05ec6562ced5788"
)

type AuthResponse map[string]string

func PerformOAuth(
	email,
	masterToken,
	gaid,
	service,
	app,
	clientSig,
	deviceCountry,
	operatorCountry,
	lang string,
	sdkVersion int) (AuthResponse, error) {
	data := url.Values{
		"accountType":     []string{"HOSTED_OR_GOOGLE"},
		"Email":           []string{email},
		"has_permission":  []string{"1"},
		"EncryptedPasswd": []string{masterToken},
		"service":         []string{service},
		"source":          []string{"android"},
		"androidId":       []string{gaid},
		"app":             []string{app},
		"client_sig":      []string{clientSig},
		"device_country":  []string{deviceCountry},
		"operatorCountry": []string{operatorCountry},
		"lang":            []string{lang},
		"sdk_version":     []string{fmt.Sprint(sdkVersion)},
	}
    return authRequest(data)
}

func PerformOAuthWithDefaults(
	email,
	masterToken,
	gaid,
	service,
	app string) (AuthResponse, error) {
	return PerformOAuth(email,
		masterToken,
		gaid,
		service,
		app,
		DefaultClientSig, "us", "us", "en", 17)
}

func authRequest(data url.Values) (AuthResponse, error) {
	req, err := http.NewRequest(http.MethodPost, authURL, strings.NewReader(data.Encode()))
	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("gpsoauth: %v", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gpsoauth: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gpsoauth: %s: %s", resp.Status, b)
	}

    return parseAuthResponse(b), nil
}

func parseAuthResponse(response []byte) AuthResponse {
    result := AuthResponse{}
	for _, line := range strings.Split(string(response), "\n") {
		sp := strings.SplitN(line, "=", 2)
		if len(sp) != 2 {
			continue
		}
        result[sp[0]] = sp[1]
	}
    return result
}
