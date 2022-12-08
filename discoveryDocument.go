package qboauth

import "time"

type Document struct {
	environment                       Environment
	refreshedAt                       time.Time
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	ClaimsSupported                   []string `json:"claims_supported"`
	IdTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	Issuer                            string   `json:"issuer"`
	JwksUri                           string   `json:"jwks_uri"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	RevocationEndpoint                string   `json:"revocation_endpoint"`
	ScopesSupported                   []string `json:"scopes_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
}

var doc *Document

type jsonGetter interface {
	getFromUrl() (Document, error)
}

func get(c jsonGetter, env Environment) (Document, error) {
	if doc == nil {
		doc = new(Document)
	}
	if !isOutdated(doc) && doc.isFor(env) {
		return *doc, nil
	}
	return c.getFromUrl()
}

func (c Client) GetDiscoveryDocument() (Document, error) {
	d, err := get(c, c.env)
	if err != nil {
		return Document{}, err
	}
	doc = &d
	return d, nil
}
