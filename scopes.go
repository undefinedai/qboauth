package qboauth

type Scope int64

const (
	Accounting Scope = iota
	Address
	Email
	OpenID
	Payments
	Phone
	Profile
)

func shouldAddOpenID(scopes []string) bool {
	contains := false
	requires := false
	for _, v := range scopes {
		if v == "openid" {
			contains = true
		}
		if v == "address" || v == "email" || v == "phone" || v == "profile" {
			requires = true
		}
	}
	return requires && !contains
}

func scopesToStrings(scopes []Scope) []string {
	var s []string
	for _, v := range scopes {
		switch v {
		case Accounting:
			s = append(s, "com.intuit.quickbooks.accounting")
		case Address:
			s = append(s, "address")
		case Email:
			s = append(s, "email")
		case OpenID:
			s = append(s, "openid")
		case Payments:
			s = append(s, "com.intuit.quickbooks.payment")
		case Phone:
			s = append(s, "phone")
		case Profile:
			s = append(s, "profile")
		}
	}
	if shouldAddOpenID(s) {
		s = append(s, "openid")
	}
	return s
}
