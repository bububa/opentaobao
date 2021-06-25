package model

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

type File struct {
	io.Reader
	Name string
}

func (f File) MultipartFileName() string {
	return fmt.Sprintf("@%s", f.Name)
}

type Param struct {
	str string
	fd  *File
}

func NewStringParam(str string) *Param {
	return &Param{
		str: str,
	}
}

func NewFileParam(fd *File) *Param {
	return &Param{
		fd: fd,
	}
}

func (p Param) IsFile() bool {
	return p.fd != nil
}

func (p Param) File() *File {
	return p.fd
}

func (p Param) String() string {
	if p.IsFile() {
		return p.File().MultipartFileName()
	}
	return p.str
}

type Params map[string]*Param

func NewParams() Params {
	p := make(Params)
	return p
}

func (p Params) Set(key string, value interface{}) error {
	switch value.(type) {
	case *File:
		p[key] = NewFileParam(value.(*File))
	default:
		str, err := AnyToString(value)
		if err != nil {
			return err
		}
		p[key] = NewStringParam(str)
	}
	return nil
}

func (p Params) GetRawParams() map[string]*Param {
	return p
}

func (p Params) NeedMultipart() bool {
	for _, v := range p {
		if v.IsFile() {
			return true
		}
	}
	return false
}

func AnyToString(val interface{}) (string, error) {
	var str string
	switch val.(type) {
	case string:
		str = val.(string)
	case int64:
		str = strconv.FormatInt(val.(int64), 10)
	case bool:
		str = "false"
		if val.(bool) {
			str = "true"
		}
	case float64:
		str = strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case *File:
		if val == nil {
			return "file", nil
		}
		return val.(*File).MultipartFileName(), nil
	default:
		data, err := json.Marshal(val)
		if err != nil {
			return str, err
		}
		str = string(data)
	}
	return str, nil
}
