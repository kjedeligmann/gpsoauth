# Ad hoc Go client package for Google Play Services OAuth

Inspired by the [gpsoauth](https://github.com/simon-weber/gpsoauth) Python library.

## Usage

For now, you can use this package to fetch an access token for authentication (that's my usecase) as follows:

```go
	accessToken, err := gpsoauth.performOAuthWithDefaults(email, masterToken, gaid, scopes, "com.google.android.keep")
```
