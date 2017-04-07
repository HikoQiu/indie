package core

import (
    "reflect"
    "errors"
)

const (
    MODULE_FIELD_NAME string = "core.Module"
)

type Module struct {
    Name string //  module name, for reading configure field
}

// check whether inherit from Module
func inheritModule(m interface{}) (bool, error) {
    flag := false
    var err error

    r := reflect.ValueOf(m)
    for i := 0; i < r.NumField(); i++ {
	f := r.Field(i)

	switch f.Type().String() {
	case MODULE_FIELD_NAME:
	    flag = true
	}
    }

    if !flag {
	err = errors.New(reflect.TypeOf(m).String() + " did not inherit from " + MODULE_FIELD_NAME)
    }

    return flag, err
}
