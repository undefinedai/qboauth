# Quickbooks OAuth Utility

A Golang utility package to help with the Quickbooks OAuth flow for custom Quickbooks apps.

[See the Quickbooks developer documentation](https://developer.intuit.com/app/developer/qbo/docs/develop/authentication-and-authorization) to learn more about how OAuth is implemented for Quickbooks.

**Note:** Undefined is not associated with Quickbooks in any way and offers no support for this module or for Quickbooks in general.

**Note:** This package is very much in alpha and should not be used in production as it is likely to change in breaking ways until it reaches version 1.0.0.

## Installation

```bash
go get github.com/undefinedai/qboauth
```

## Usage

### Quickbooks OAuth Client

```golang
client := qboauth.NewClient(...)
```

`NewClient` takes the following arguments:

- `clientId` (string) - the Client ID for the Quickbooks app
- `clientSecret` (string) - the Client Secret for the Quickbooks app
- `redirectURL` (string) - the url to which Quickbooks should redirect after a user has successfully authenticated with Quickbooks and connected the app
- `scopes` ([]qboauth.Scope) - an array of Scopes for the Quickbooks app (one or more of: `Accounting`, `Address`, `Email`, `OpenID`, `Payments`, `Phone`, `Profile`)
- `env` (qboauth.Environment) - either `Sandbox` or `Production`

### Connecting your app to a Quickbooks company

The first step is to send your user to Quickbooks the url retrieved from `client.GetAuthCodeURL` to begin the authentication flow. You must include a `state` string which you can later use to validate the response from Quickbooks. If you do not have a specific `state` that you need to track for your user, this can be a uuid or any string, keeping in mind that it will be appended to the url as a query parameter and as such, will be subject to the maximum length of allowable urls.

You'll receive this `state` back after Quickbooks has authenticated the user and the user has connected your app to their Quickbooks company.

```golang
url, err := client.GetAuthCodeURL(state)
```

Quickbooks will authenticate the user and redirect to the `redirectURL` you provided in `NewClient`. You'll receive (at least) the `state` you provided, a `realmId` (the Quickbooks company id), and a `code` as url query parameters when the user is sent back to your `redirectURL`. Verify that the `state` is correct, note or store the `realmId`, and pass the `code` to:

```golang
tokens, err := c.Exchange(code)
```

You'll receive a `qboauth.Tokens`:

```golang
type Tokens struct {
	AccessToken                   string
	ExpiresInSeconds              int
	IDToken                       string
	RefreshToken                  string
	TokenType                     string
	XRefreshTokenExpiresInSeconds int
}
```

**Note:** `IDToken` will only be present if an openid scope is requested.

### Using, Refreshing, and Revoking Tokens

The `AccessToken` and `IDToken` are only valid for one hour - you will know how long they're valid by checking `ExpiresInSeconds`. When (or before) they expire, you can use the `RefreshToken` to generate updated tokens:

```golang
tokens, err := c.Refresh(refreshToken)
```

It's important to note that (a) the `RefreshToken` is only valid for 100 days and (b) may change from time to time; so, each time you refresh the tokens, note the new `RefreshToken` and use it in future requests so that your app remains connected. If the `RefreshToken` expires (you can track this with `XRefreshTokenExpiresInSeconds`) or is otherwise revoked, your user will have to go through the authentication process again.

To intentionally invalidate/revoke a set of tokens (such as when a client uninstalls your app), you can pass the `RefreshToken` to:

```golang
err := c.Revoke(refreshToken)
```

### Quickbooks Discovery Document

If you're verifying or creating JWTs to be used with the Quickbooks api, you may want to access the discovery document directly. This package retrieves the discovery document and stores it in memory for up to 24 hours to reduce http requests and latency on subsequent requests, and can access it directly with:

```golang
doc, err := c.GetDiscoveryDocument()
```

The discovery document retrieved is based on the `Environment` you provide to `NewClient`.
