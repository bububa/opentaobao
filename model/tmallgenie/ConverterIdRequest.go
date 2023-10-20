package tmallgenie

import (
	"sync"
)

// ConverterIdRequest 结构体
type ConverterIdRequest struct {
	// idType相应的id内容
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 编码类型，此处填写PROJECT_ID。
	EncodeType string `json:"encode_type,omitempty" xml:"encode_type,omitempty"`
	// 编码类型对应的值，此处填写该产品所在项目的Project ID。请在天猫精灵AI平台的控制台中查看。
	EncodeKey string `json:"encode_key,omitempty" xml:"encode_key,omitempty"`
	// USER_ID/DEVICE_ID/OPEN_TAOBAO_ID
	IdType string `json:"id_type,omitempty" xml:"id_type,omitempty"`
}

var poolConverterIdRequest = sync.Pool{
	New: func() any {
		return new(ConverterIdRequest)
	},
}

// GetConverterIdRequest() 从对象池中获取ConverterIdRequest
func GetConverterIdRequest() *ConverterIdRequest {
	return poolConverterIdRequest.Get().(*ConverterIdRequest)
}

// ReleaseConverterIdRequest 释放ConverterIdRequest
func ReleaseConverterIdRequest(v *ConverterIdRequest) {
	v.Id = ""
	v.EncodeType = ""
	v.EncodeKey = ""
	v.IdType = ""
	poolConverterIdRequest.Put(v)
}
