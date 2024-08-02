# Ad hoc Go client package for Google Play Services OAuth

Inspired by the [gpsoauth](https://github.com/simon-weber/gpsoauth) Python library.

## Usage

For now, you can use this package to fetch an access token for authentication (that's my usecase) as follows:

```go
package main

import (
	"fmt"
	"log"

	"github.com/kjedeligmann/gpsoauth"
)

func main() {
	var email, masterToken, gaid, scopes string // you should have credentials and API scopes

	resp, err := gpsoauth.PerformOAuthWithDefaults(
		email,
		masterToken,
		gaid, // Google Advertising ID, or Android ID
		scopes,
		"com.google.android.keep")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp["Auth"]) // access token
}
```

This is similar to [this piece of code](https://github.com/kiwiz/gkeepapi/blob/d56a9e388dc66a51ec7d0dd37e443c2beb37f5a7/src/gkeepapi/__init__.py#L149C1-L163C39).

You can also fetch a master token for your account:

```go
	resp, err := gpsoauth.ExchangeTokenWithDefaults(email, webToken, gaid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp["Token"]) // master token
```

You can learn more about obtaining `webToken` [here](https://github.com/rukins/gpsoauth-java/blob/b74ebca999d0f5bd38a2eafe3c0d50be552f6385/README.md#receiving-an-authentication-token). Overview of Google Play Services OAuth flow can be found [here](https://sbktech.blogspot.com/2014/01/inside-android-play-services-magic.html). 
