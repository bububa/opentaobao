package damai

import (
	"sync"
)

// ThirdTicketFacePushOpenParam 结构体
type ThirdTicketFacePushOpenParam struct {
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 票面id
	FaceId int64 `json:"face_id,omitempty" xml:"face_id,omitempty"`
	// 票面模式
	FaceMode int64 `json:"face_mode,omitempty" xml:"face_mode,omitempty"`
	// 票面类型
	FrontType int64 `json:"front_type,omitempty" xml:"front_type,omitempty"`
	// 票纸版式id
	PaperFormatId int64 `json:"paper_format_id,omitempty" xml:"paper_format_id,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}

var poolThirdTicketFacePushOpenParam = sync.Pool{
	New: func() any {
		return new(ThirdTicketFacePushOpenParam)
	},
}

// GetThirdTicketFacePushOpenParam() 从对象池中获取ThirdTicketFacePushOpenParam
func GetThirdTicketFacePushOpenParam() *ThirdTicketFacePushOpenParam {
	return poolThirdTicketFacePushOpenParam.Get().(*ThirdTicketFacePushOpenParam)
}

// ReleaseThirdTicketFacePushOpenParam 释放ThirdTicketFacePushOpenParam
func ReleaseThirdTicketFacePushOpenParam(v *ThirdTicketFacePushOpenParam) {
	v.PushTime = ""
	v.SupplierSecret = ""
	v.FaceId = 0
	v.FaceMode = 0
	v.FrontType = 0
	v.PaperFormatId = 0
	v.PerformId = 0
	v.SystemId = 0
	poolThirdTicketFacePushOpenParam.Put(v)
}
