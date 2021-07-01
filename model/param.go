package model

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

// 用于上传文件，对应淘宝API中byte[]类型
type File struct {
	io.Reader
	Name string
}

func (f File) MultipartFileName() string {
	return fmt.Sprintf("@%s", f.Name)
}

// APIRequest中的参数值
type ParamValue struct {
	str string
	fd  *File
}

// string类型值
func NewStringParamValue(str string) *ParamValue {
	return &ParamValue{
		str: str,
	}
}

// file类型值
func NewFileParamValue(fd *File) *ParamValue {
	return &ParamValue{
		fd: fd,
	}
}

func (p ParamValue) IsFile() bool {
	return p.fd != nil
}

func (p ParamValue) File() *File {
	return p.fd
}

func (p ParamValue) String() string {
	if p.IsFile() {
		return p.File().MultipartFileName()
	}
	return p.str
}

// APIRequest中的参数列表
type Params map[string]*ParamValue

func NewParams() Params {
	p := make(Params)
	return p
}

// APIRequest设置参数
func (p Params) Set(key string, value interface{}) error {
	switch value.(type) {
	case *File:
		p[key] = NewFileParamValue(value.(*File))
	default:
		str, err := AnyToString(value)
		if err != nil {
			return err
		}
		p[key] = NewStringParamValue(str)
	}
	return nil
}

func (p Params) GetRawParams() map[string]*ParamValue {
	return p
}

// 是否需要使用form/multipart post
func (p Params) NeedMultipart() bool {
	for _, v := range p {
		if v.IsFile() {
			return true
		}
	}
	return false
}

// 转换任意类型参数值为string
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
