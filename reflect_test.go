// Copyright 2021 projectred. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tinject

import (
	"reflect"
	"testing"
)

type A interface {
	Value() int
}

type DefaultA struct {
	V int
}

func (a *DefaultA) Value() int { return a.V }

func TestNewStruct(t *testing.T) {
	if err := Regist(reflect.TypeOf((*DefaultA)(nil)), RegistOptionName("defaultA")); err != nil {
		t.Error(err)
		return
	}
	println(RegistList())
	var a A = NewStructByKeyName("github.com/projectred/tinject.DefaultA", NewStructKvs(KV{"V", 15})).(A)
	if a.Value() != 15 {
		t.Errorf("it should be %d, but it is %d", 15, a.Value())
	}
	a = NewStructByKeyName("defaultA", NewStructKvs(KV{"V", 17})).(A)
	if a.Value() != 17 {
		t.Errorf("it should be %d, but it is %d", 17, a.Value())
	}
	Fill(reflect.ValueOf(a), []KV{{K: "V", V: 20}})
	if a.Value() != 20 {
		t.Errorf("it should be %d, but it is %d", 20, a.Value())
	}
}
