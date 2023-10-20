package damai

import (
	"sync"
)

// ThirdPaperFormatPushOpenParam 结构体
type ThirdPaperFormatPushOpenParam struct {
	// 票纸版式名称
	PaperFormatName string `json:"paper_format_name,omitempty" xml:"paper_format_name,omitempty"`
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 票纸版式id
	PaperFormatId int64 `json:"paper_format_id,omitempty" xml:"paper_format_id,omitempty"`
	// 打印类型
	PrintType int64 `json:"print_type,omitempty" xml:"print_type,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}

var poolThirdPaperFormatPushOpenParam = sync.Pool{
	New: func() any {
		return new(ThirdPaperFormatPushOpenParam)
	},
}

// GetThirdPaperFormatPushOpenParam() 从对象池中获取ThirdPaperFormatPushOpenParam
func GetThirdPaperFormatPushOpenParam() *ThirdPaperFormatPushOpenParam {
	return poolThirdPaperFormatPushOpenParam.Get().(*ThirdPaperFormatPushOpenParam)
}

// ReleaseThirdPaperFormatPushOpenParam 释放ThirdPaperFormatPushOpenParam
func ReleaseThirdPaperFormatPushOpenParam(v *ThirdPaperFormatPushOpenParam) {
	v.PaperFormatName = ""
	v.PushTime = ""
	v.SupplierSecret = ""
	v.PaperFormatId = 0
	v.PrintType = 0
	v.SystemId = 0
	poolThirdPaperFormatPushOpenParam.Put(v)
}
