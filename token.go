package qboauth

type Tokens struct {
	AccessToken                   string `json:"access_token,omitempty"`
	ExpiresInSeconds              int    `json:"expires_in,omitempty"`
	IDToken                       string `json:"id_token,omitempty"`
	RefreshToken                  string `json:"refresh_token,omitempty"`
	TokenType                     string `json:"token_type,omitempty"`
	XRefreshTokenExpiresInSeconds int    `json:"x_refresh_token_expires_in,omitempty"`
}
