package xhotelitem

import (
	"sync"
)

// ServiceTimeDataDo 结构体
type ServiceTimeDataDo struct {
	// supplier
	Supplier string `json:"supplier,omitempty" xml:"supplier,omitempty"`
	// 卖家nick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 业务类型：1卖家；2supplier;3酒店
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// timeZoneName
	TimeZoneName string `json:"time_zone_name,omitempty" xml:"time_zone_name,omitempty"`
	// 周五服务时间（当地时间）
	FridayConfirmLocalTime string `json:"friday_confirm_local_time,omitempty" xml:"friday_confirm_local_time,omitempty"`
	// 周一服务时间（当地时间）
	MondayConfirmLocalTime string `json:"monday_confirm_local_time,omitempty" xml:"monday_confirm_local_time,omitempty"`
	// 周二服务时间（当地时间）
	TuesdayConfirmLocalTime string `json:"tuesday_confirm_local_time,omitempty" xml:"tuesday_confirm_local_time,omitempty"`
	// 周三服务时间（当地时间）
	WednesdayConfirmLocalTime string `json:"wednesday_confirm_local_time,omitempty" xml:"wednesday_confirm_local_time,omitempty"`
	// 周六服务时间（当地时间）
	SaturdayConfirmLocalTime string `json:"saturday_confirm_local_time,omitempty" xml:"saturday_confirm_local_time,omitempty"`
	// operator
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 周日服务时间（当地时间）
	SundayConfirmLocalTime string `json:"sunday_confirm_local_time,omitempty" xml:"sunday_confirm_local_time,omitempty"`
	// 周四服务时间（当地时间）
	ThursdayConfirmLocalTime string `json:"thursday_confirm_local_time,omitempty" xml:"thursday_confirm_local_time,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 最后修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 业务id
	BusinessId int64 `json:"business_id,omitempty" xml:"business_id,omitempty"`
	// 1: 当日订单, 2:次日及以后订单
	OrderConfirmType int64 `json:"order_confirm_type,omitempty" xml:"order_confirm_type,omitempty"`
	// 是否在非工作时间显示商品 1:显示, 2:不显示
	DisplayItemInNonworkingTime int64 `json:"display_item_in_nonworking_time,omitempty" xml:"display_item_in_nonworking_time,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolServiceTimeDataDo = sync.Pool{
	New: func() any {
		return new(ServiceTimeDataDo)
	},
}

// GetServiceTimeDataDo() 从对象池中获取ServiceTimeDataDo
func GetServiceTimeDataDo() *ServiceTimeDataDo {
	return poolServiceTimeDataDo.Get().(*ServiceTimeDataDo)
}

// ReleaseServiceTimeDataDo 释放ServiceTimeDataDo
func ReleaseServiceTimeDataDo(v *ServiceTimeDataDo) {
	v.Supplier = ""
	v.SellerNick = ""
	v.BusinessType = ""
	v.TimeZoneName = ""
	v.FridayConfirmLocalTime = ""
	v.MondayConfirmLocalTime = ""
	v.TuesdayConfirmLocalTime = ""
	v.WednesdayConfirmLocalTime = ""
	v.SaturdayConfirmLocalTime = ""
	v.Operator = ""
	v.SundayConfirmLocalTime = ""
	v.ThursdayConfirmLocalTime = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.SellerId = 0
	v.BusinessId = 0
	v.OrderConfirmType = 0
	v.DisplayItemInNonworkingTime = 0
	v.Id = 0
	poolServiceTimeDataDo.Put(v)
}
