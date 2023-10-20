package damai

import (
	"sync"
)

// TicketFaceIdOpenParam 结构体
type TicketFaceIdOpenParam struct {
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 票面id
	FaceId int64 `json:"face_id,omitempty" xml:"face_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}

var poolTicketFaceIdOpenParam = sync.Pool{
	New: func() any {
		return new(TicketFaceIdOpenParam)
	},
}

// GetTicketFaceIdOpenParam() 从对象池中获取TicketFaceIdOpenParam
func GetTicketFaceIdOpenParam() *TicketFaceIdOpenParam {
	return poolTicketFaceIdOpenParam.Get().(*TicketFaceIdOpenParam)
}

// ReleaseTicketFaceIdOpenParam 释放TicketFaceIdOpenParam
func ReleaseTicketFaceIdOpenParam(v *TicketFaceIdOpenParam) {
	v.SupplierSecret = ""
	v.FaceId = 0
	v.SystemId = 0
	poolTicketFaceIdOpenParam.Put(v)
}
