// Copyright 2021 projectred. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tinject

import (
	"reflect"
)

type KV struct {
	K string
	V interface{}
}
type NewStructOption struct {
	kvs []KV
}

type NewStructF func(*NewStructOption)

func NewStructKvs(kvs ...KV) NewStructF {
	return func(o *NewStructOption) { o.kvs = kvs }
}

func NewStruct(t reflect.Type, ofs ...NewStructF) interface{} {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	result := reflect.New(t)
	var o NewStructOption
	for _, f := range ofs {
		f(&o)
	}
	fill(result, o.kvs)
	return result.Interface()
}

func fill(v reflect.Value, kvs []KV) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for _, kv := range kvs {
		if field := v.FieldByName(kv.K); field.CanSet() {
			field.Set(reflect.ValueOf(kv.V))
		}
	}
}
