package wdk

import (
	"sync"
)

// Postman 结构体
type Postman struct {
	// 配送员姓名
	PostmanName string `json:"postman_name,omitempty" xml:"postman_name,omitempty"`
	// 配送员编码
	PostmanCode string `json:"postman_code,omitempty" xml:"postman_code,omitempty"`
	// 配送员手机
	PostmanPhone string `json:"postman_phone,omitempty" xml:"postman_phone,omitempty"`
	// 服务商名称
	ProviderName string `json:"provider_name,omitempty" xml:"provider_name,omitempty"`
	// 服务商编码
	ProviderCode string `json:"provider_code,omitempty" xml:"provider_code,omitempty"`
}

var poolPostman = sync.Pool{
	New: func() any {
		return new(Postman)
	},
}

// GetPostman() 从对象池中获取Postman
func GetPostman() *Postman {
	return poolPostman.Get().(*Postman)
}

// ReleasePostman 释放Postman
func ReleasePostman(v *Postman) {
	v.PostmanName = ""
	v.PostmanCode = ""
	v.PostmanPhone = ""
	v.ProviderName = ""
	v.ProviderCode = ""
	poolPostman.Put(v)
}
