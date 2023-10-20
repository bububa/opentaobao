package damai

import (
	"sync"
)

// VenueIdOpenParam 结构体
type VenueIdOpenParam struct {
	// 操作员
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 商家密钥，监权使用
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 商家id
	CpId int64 `json:"cp_id,omitempty" xml:"cp_id,omitempty"`
	// 推送接口返回的id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 操作员id
	OperatorId int64 `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 商家id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 场馆id
	VenueId int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
}

var poolVenueIdOpenParam = sync.Pool{
	New: func() any {
		return new(VenueIdOpenParam)
	},
}

// GetVenueIdOpenParam() 从对象池中获取VenueIdOpenParam
func GetVenueIdOpenParam() *VenueIdOpenParam {
	return poolVenueIdOpenParam.Get().(*VenueIdOpenParam)
}

// ReleaseVenueIdOpenParam 释放VenueIdOpenParam
func ReleaseVenueIdOpenParam(v *VenueIdOpenParam) {
	v.Operator = ""
	v.SupplierSecret = ""
	v.CpId = 0
	v.Id = 0
	v.OperatorId = 0
	v.SystemId = 0
	v.VenueId = 0
	poolVenueIdOpenParam.Put(v)
}
