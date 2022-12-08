package qboauth

func urlFor(env Environment) string {
	if env == Production {
		return "https://developer.api.intuit.com/.well-known/openid_configuration"
	}
	return "https://developer.api.intuit.com/.well-known/openid_sandbox_configuration"
}
