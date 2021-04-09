package policies

import (
	"os"
)

var basePath = os.Getenv("POLICY_PATH")

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
	return basePath + string(p) + "/" + string(p) + ".conf"
}

func (p policy) GetPolicyCsvPath() string {
	return basePath + string(p) + "/" + string(p) + ".csv"
}
