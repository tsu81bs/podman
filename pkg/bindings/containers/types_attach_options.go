package containers

import (
	"net/url"
	"reflect"
	"strings"

	"github.com/containers/podman/v2/pkg/bindings/util"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

/*
This file is generated automatically by go generate.  Do not edit.
*/

// Changed
func (o *AttachOptions) Changed(fieldName string) bool {
	r := reflect.ValueOf(o)
	value := reflect.Indirect(r).FieldByName(fieldName)
	return !value.IsNil()
}

// ToParams
func (o *AttachOptions) ToParams() (url.Values, error) {
	params := url.Values{}
	if o == nil {
		return params, nil
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	s := reflect.ValueOf(o)
	if reflect.Ptr == s.Kind() {
		s = s.Elem()
	}
	sType := s.Type()
	for i := 0; i < s.NumField(); i++ {
		fieldName := sType.Field(i).Name
		if !o.Changed(fieldName) {
			continue
		}
		fieldName = strings.ToLower(fieldName)
		f := s.Field(i)
		if reflect.Ptr == f.Kind() {
			f = f.Elem()
		}
		switch {
		case util.IsSimpleType(f):
			params.Set(fieldName, util.SimpleTypeToParam(f))
		case f.Kind() == reflect.Slice:
			for i := 0; i < f.Len(); i++ {
				elem := f.Index(i)
				if util.IsSimpleType(elem) {
					params.Add(fieldName, util.SimpleTypeToParam(elem))
				} else {
					return nil, errors.New("slices must contain only simple types")
				}
			}
		case f.Kind() == reflect.Map:
			lowerCaseKeys := make(map[string][]string)
			iter := f.MapRange()
			for iter.Next() {
				lowerCaseKeys[iter.Key().Interface().(string)] = iter.Value().Interface().([]string)

			}
			s, err := json.MarshalToString(lowerCaseKeys)
			if err != nil {
				return nil, err
			}

			params.Set(fieldName, s)
		}

	}
	return params, nil
}

// WithDetachKeys
func (o *AttachOptions) WithDetachKeys(value string) *AttachOptions {
	v := &value
	o.DetachKeys = v
	return o
}

// GetDetachKeys
func (o *AttachOptions) GetDetachKeys() string {
	var detachKeys string
	if o.DetachKeys == nil {
		return detachKeys
	}
	return *o.DetachKeys
}

// WithLogs
func (o *AttachOptions) WithLogs(value bool) *AttachOptions {
	v := &value
	o.Logs = v
	return o
}

// GetLogs
func (o *AttachOptions) GetLogs() bool {
	var logs bool
	if o.Logs == nil {
		return logs
	}
	return *o.Logs
}

// WithStream
func (o *AttachOptions) WithStream(value bool) *AttachOptions {
	v := &value
	o.Stream = v
	return o
}

// GetStream
func (o *AttachOptions) GetStream() bool {
	var stream bool
	if o.Stream == nil {
		return stream
	}
	return *o.Stream
}