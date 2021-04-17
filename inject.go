// Copyright 2021 projectred. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tinject

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

type Injects struct {
	Default map[string]reflect.Type
	Alias   map[string]reflect.Type
}

var defaultInject = Injects{Default: make(map[string]reflect.Type), Alias: make(map[string]reflect.Type)}
var Regist = defaultInject.Regist
var RegistList = defaultInject.String
var NewStructByKeyName = defaultInject.NewStruct

type RegistOption struct {
	Name *string
}

func (o RegistOption) Load(ofs ...RegistOptionF) *RegistOption {
	for _, f := range ofs {
		f(&o)
	}
	return &o
}

type RegistOptionF func(*RegistOption)

func RegistOptionName(name string) RegistOptionF {
	return func(o *RegistOption) { o.Name = &name }
}

func (i Injects) Regist(t reflect.Type, ofs ...RegistOptionF) error {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	setF := func(m map[string]reflect.Type, name string) error {
		if _, ok := m[name]; ok {
			return errors.New("type has been exist")
		}
		m[name] = t
		return nil
	}
	if err := setF(i.Default, t.PkgPath()+"."+t.Name()); err != nil {
		return err
	}
	if name := (RegistOption{}.Load(ofs...)).Name; name != nil {
		if err := setF(i.Alias, *name); err != nil {
			return err
		}
	}
	return nil
}

func (i Injects) String() string {
	result := bytes.NewBufferString("")
	for _, inject := range []struct {
		token string
		m     map[string]reflect.Type
	}{
		{"pkg.type.Name", i.Default},
		{"alias", i.Alias},
	} {
		for k, v := range inject.m {
			result.WriteString(fmt.Sprintf("%s: %s : %s\n", inject.token, k, v.Name()))
		}
	}
	return result.String()
}

func (i Injects) NewStruct(key string, ofs ...NewStructF) interface{} {
	for _, m := range []map[string]reflect.Type{i.Alias, i.Default} {
		if t, ok := m[key]; ok {
			return NewStruct(t, ofs...)
		}
	}
	return nil
}
