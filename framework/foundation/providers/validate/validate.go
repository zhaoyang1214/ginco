package validate

import (
	"ginco/framework/contract"
	"github.com/go-playground/validator/v10"
)

type Validate struct {
}

var _ contract.Provider = (*Validate)(nil)

func (v *Validate) Build(container contract.Container, params ...interface{}) (interface{}, error) {
	return validator.New(), nil
}
