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
	p, admin,/api/v1/users/:user_id/images,PUT
	p, admin, /api/v1/users/:user_id, DELETE
	p, admin, /api/v1/users/invitations, POST
	p, admin, /api/v1/users/:user_id/permissions/update, PUT
	p, admin, /api/v1/invitations/:invitation_id, DELETE
	p, admin, /api/v1/invitations/:invitation_id/resend, POST
	p, admin, /api/v1/retailers/:retailer_id/google-my-business/*, POST
	p, hp, /api/v1/omnichat/*, GET
	p, hp, /api/v1/pc-selector/*, GET
	p, hp, /api/v1/sentiment/*, GET
	p, hp, /api/v1/indexed/*, GET
	p, hp,/api/v1/countries, GET
	p, hp, /api/v1/countries/*, GET
	p, hp,/api/v1/latam/*, GET
	p, hp, /api/v1/retailers, GET
	p, hp, /api/v1/retailers/*, GET
	p, hp, /api/v1/sectors, GET
	p, hp, /api/v1/categories, GET
	p, hp, /api/v1/multichannel/*, GET
	p, hp, /api/v1/retailers/:retailer_id/google-my-business/*, POST
	p, hp,/api/v1/users/:user_id/images,PUT
	p, hp, /api/v1/activeCountries, GET
	p, country, /api/v1/omnichat/*,GET
	p, country, /api/v1/pc-selector/*,GET
	p, country, /api/v1/indexed/*,GET
	p, country,/api/v1/latam/*, GET
	p, country,/api/v1/countries, GET
	p, country, /api/v1/countries/*, GET
	p, country, /api/v1/sectors, GET
	p, country, /api/v1/categories, GET
	p, country, /api/v1/retailers, GET
	p, country, /api/v1/retailers/*, GET
	p, country, /api/v1/multichannel/*, GET
	p, country, /api/v1/retailers/:retailer_id/google-my-business/*, POST
	p, country,/api/v1/users/:user_id/images,PUT
	p, country, /api/v1/activeCountries, GET
	p, retailer, /api/v1/omnichat/retailers/*, GET
	p, retailer, /api/v1/pc-selector/retailers/*, GET
	p, retailer, /api/v1/indexed/retailers/*, GET
	p, retailer, /api/v1/retailers, GET
	p, retailer, /api/v1/retailers/*, GET
	p, retailer, /api/v1/sectors, GET
	p, retailer, /api/v1/categories, GET
	p, retailer, /api/v1/multichannel/retailers/*, GET
	p, retailer, /api/v1/retailers/:retailer_id/google-my-business/*, POST
	p, retailer,/api/v1/users/:user_id/images,PUT
	p, retailer, /api/v1/activeCountries, GET
	`
	_, err = file.Write([]byte(d2))
	if err != nil {
		return ""
	}

	return fileName
}
