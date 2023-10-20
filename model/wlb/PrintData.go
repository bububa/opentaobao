package wlb

import (
	"sync"
)

// PrintData 结构体
type PrintData struct {
	// 版本
	Ver string `json:"ver,omitempty" xml:"ver,omitempty"`
	//  打印数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 签名
	Signature string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 追加的 data
	AddData string `json:"add_data,omitempty" xml:"add_data,omitempty"`
	//  模板 url
	TemplateUrl string `json:"template_url,omitempty" xml:"template_url,omitempty"`
	//  是否加密
	Encrypted bool `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
}

var poolPrintData = sync.Pool{
	New: func() any {
		return new(PrintData)
	},
}

// GetPrintData() 从对象池中获取PrintData
func GetPrintData() *PrintData {
	return poolPrintData.Get().(*PrintData)
}

// ReleasePrintData 释放PrintData
func ReleasePrintData(v *PrintData) {
	v.Ver = ""
	v.Data = ""
	v.Signature = ""
	v.AddData = ""
	v.TemplateUrl = ""
	v.Encrypted = false
	poolPrintData.Put(v)
}
