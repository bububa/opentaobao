package ascp

import (
	"sync"
)

// SpecialWarehouseTimingRules 结构体
type SpecialWarehouseTimingRules struct {
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
	// 截单时间，例如17:00，只允许整点或者半点； 当业务身份为时效代运营时，截单时间前支付订单，承诺今日发，否则承诺24小时发
	CutTime string `json:"cut_time,omitempty" xml:"cut_time,omitempty"`
	// 过截单时间后，不表达时间段开始时间（HH）
	BanHourFrom string `json:"ban_hour_from,omitempty" xml:"ban_hour_from,omitempty"`
	// 过截单时间后，不表达时间段结束时间（HH） ● banHourFrom、banHourTo需要一起出现；相隔时间≤5小时，只允许整点
	BanHourTo string `json:"ban_hour_to,omitempty" xml:"ban_hour_to,omitempty"`
	// 最晚接单截单时间（HH:mm），例如17:00，只允许整点或者半点
	ReceiveCutTime string `json:"receive_cut_time,omitempty" xml:"receive_cut_time,omitempty"`
	// 是否承诺发 1=承诺发；当截单时间前支付/截单的订单，承诺今日发，否则，承诺24小时发 0=不承诺
	PromiseType int64 `json:"promise_type,omitempty" xml:"promise_type,omitempty"`
}

var poolSpecialWarehouseTimingRules = sync.Pool{
	New: func() any {
		return new(SpecialWarehouseTimingRules)
	},
}

// GetSpecialWarehouseTimingRules() 从对象池中获取SpecialWarehouseTimingRules
func GetSpecialWarehouseTimingRules() *SpecialWarehouseTimingRules {
	return poolSpecialWarehouseTimingRules.Get().(*SpecialWarehouseTimingRules)
}

// ReleaseSpecialWarehouseTimingRules 释放SpecialWarehouseTimingRules
func ReleaseSpecialWarehouseTimingRules(v *SpecialWarehouseTimingRules) {
	v.WmsOwnerCode = ""
	v.CutTime = ""
	v.BanHourFrom = ""
	v.BanHourTo = ""
	v.ReceiveCutTime = ""
	v.PromiseType = 0
	poolSpecialWarehouseTimingRules.Put(v)
}
