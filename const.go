package mpc

type RunMode string

const (
	Develop    RunMode = "dev"
	Test       RunMode = "qa"
	Staging    RunMode = "pre"
	Production RunMode = "prd"
)

func (r RunMode) String() string {
	return string(r)
}

func (r RunMode) IsValid() bool {
	switch r {
	case Develop, Test, Staging, Production:
		return true
	}
	return false
}