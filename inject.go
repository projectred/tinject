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

type Injects map[string]reflect.Type

var defaultInject = make(Injects)
var Regist = defaultInject.Regist
var RegistList = defaultInject.String
var NewStructByKeyName = defaultInject.NewStruct

func (i Injects) Regist(t reflect.Type) error {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if _, ok := i[t.PkgPath()+"."+t.Name()]; ok {
		return errors.New("regist type has been exist")
	}
	i[t.PkgPath()+"."+t.Name()] = t
	return nil
}

func (i Injects) String() string {
	result := bytes.NewBufferString("")
	for k, v := range i {
		result.WriteString(fmt.Sprintf("%s : %s\n", k, v.Name()))
	}
	return result.String()
}

func (i Injects) NewStruct(t string, ofs ...NewStructF) interface{} {
	return NewStruct(i[t], ofs...)
}
