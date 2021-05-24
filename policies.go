package policies

import (
	"os"
)

type policy string

type Policy interface {
	GetPolicyConfPath() string
	GetPolicyCsvPath() string
}

const (
	TestPolicy = policy("test")
	CoopPolicy = policy("coop")
)

func (p policy) GetPolicyConfPath() string {
	fileName := "policy.conf"
	file, err := os.Create(fileName)
	if err != nil {
		return ""
	}
	d2 := `
	[request_definition]
	r = sub, obj, act
	[policy_definition]
	p = sub, obj, act
	[policy_effect]
	e = some(where (p.eft == allow))
	[matchers]
	m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
	`
	_, err = file.Write([]byte(d2))
	if err != nil {
		return ""
	}

	return fileName
}

func (p policy) GetPolicyCsvPath() string {
	fileName := "policy.csv"
	file, err := os.Create(fileName)
	if err != nil {
		return ""
	}
	d2 := `
		p, admin, *, GET
	p, admin, /api/v1/users/:user_id, DELETE
	p, admin, /api/v1/users/invitations, POST
	p, admin, /api/v1/invitations/:invitation_id, DELETE
	p, admin, /api/v1/invitations/:invitation_id/resend, POST
	p, hp,/api/v1/countries, GET
	p, hp, /api/v1/countries/*, GET
	p, hp, /api/v1/retailers/*, GET
	p, hp, /api/v1/sectors, GET
	p, hp, /api/v1/categories, GET
	p, country,/api/v1/countries, GET
	p, country, /api/v1/countries/*, GET
	p, country, /api/v1/sectors, GET
	p, country, /api/v1/categories, GET
	p, country, /api/v1/retailers, GET
	p, country, /api/v1/retailers/*, GET
	p, retailer, /api/v1/retailers, GET
	p, retailer, /api/v1/retailers/*, GET
	p, retailer, /api/v1/sectors, GET
	p, retailer, /api/v1/categories, GET
	`
	_, err = file.Write([]byte(d2))
	if err != nil {
		return ""
	}

	return fileName
}
