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
	authURL   = "https://android.clients.google.com/auth"
	userAgent = "GoogleAuth/1.4"
)

var (
	defaultClientSig = "38918a453d07199354f8b19af05ec6562ced5788"
)

func performOAuth(
	email,
	masterToken,
	gaid,
	service,
	app,
	clientSig,
	deviceCountry,
	operatorCountry,
	lang string,
	sdkVersion int) (string, error) {
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

	req, err := http.NewRequest(http.MethodPost, authURL, strings.NewReader(data.Encode()))
	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("gpsoauth: %v", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("gpsoauth: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("gpsoauth: %s: %s", resp.Status, b)
	}

	for _, line := range strings.Split(string(b), "\n") {
		sp := strings.SplitN(line, "=", 2)
		if len(sp) != 2 {
			continue
		}
		if sp[0] == "Auth" {
			return sp[1], nil
		}
	}
	return "", fmt.Errorf("gpsoauth: no Auth found")
}

func performOAuthWithDefaults(
	email,
	masterToken,
	gaid,
	service,
	app string) (string, error) {
	return performOAuth(email,
		masterToken,
		gaid,
		service,
		app,
		defaultClientSig, "us", "us", "en", 17)
}
