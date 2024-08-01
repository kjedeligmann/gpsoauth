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
	accessToken, err := gpsoauth.PerformOAuthWithDefaults(
		email,
		masterToken,
		gaid,
		scopes,
		"com.google.android.keep", "38918a453d07199354f8b19af05ec6562ced5788")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(accessToken)
}
```
