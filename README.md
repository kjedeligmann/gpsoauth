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

	accessToken, err := gpsoauth.PerformOAuthWithDefaults(
		email,
		masterToken,
		gaid, // Google Advertising ID, or Android ID
		scopes,
		"com.google.android.keep", "38918a453d07199354f8b19af05ec6562ced5788")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(accessToken)
}
```

This is similar to [this piece of code](https://github.com/kiwiz/gkeepapi/blob/d56a9e388dc66a51ec7d0dd37e443c2beb37f5a7/src/gkeepapi/__init__.py#L149C1-L163C39).
