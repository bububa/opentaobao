package trade

import (
	"sync"
)

// OrderDeliverer 结构体
type OrderDeliverer struct {
	// 配送人员编码
	DelivererCode string `json:"deliverer_code,omitempty" xml:"deliverer_code,omitempty"`
	// 配送人员电话
	DelivererPhone string `json:"deliverer_phone,omitempty" xml:"deliverer_phone,omitempty"`
	// 配送人员姓名
	DelivererName string `json:"deliverer_name,omitempty" xml:"deliverer_name,omitempty"`
	// 拣货结束时间
	PickupEndTime string `json:"pickup_end_time,omitempty" xml:"pickup_end_time,omitempty"`
	// 拣货开始时间
	PickupStartTime string `json:"pickup_start_time,omitempty" xml:"pickup_start_time,omitempty"`
	// 批次结束时间
	BatchEndTime string `json:"batch_end_time,omitempty" xml:"batch_end_time,omitempty"`
	// 批次开始时间
	BatchStartTime string `json:"batch_start_time,omitempty" xml:"batch_start_time,omitempty"`
	// 签收时间
	SignTime string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 配送结束时间
	DispatchEndTime string `json:"dispatch_end_time,omitempty" xml:"dispatch_end_time,omitempty"`
	// 配送开始时间
	DispatchStartTime string `json:"dispatch_start_time,omitempty" xml:"dispatch_start_time,omitempty"`
	// 打包结束时间
	PackageEndTime string `json:"package_end_time,omitempty" xml:"package_end_time,omitempty"`
	// 打包开始时间
	PackageStartTime string `json:"package_start_time,omitempty" xml:"package_start_time,omitempty"`
	// 签收备注
	SignMemo string `json:"sign_memo,omitempty" xml:"sign_memo,omitempty"`
	// 收货开始时间
	DeliveryStartTime string `json:"delivery_start_time,omitempty" xml:"delivery_start_time,omitempty"`
	// 收货人
	ConsigneeName string `json:"consignee_name,omitempty" xml:"consignee_name,omitempty"`
	// 收货结束时间
	DeliveryEndTime string `json:"delivery_end_time,omitempty" xml:"delivery_end_time,omitempty"`
	// 配送坐标
	DeliveryGeo string `json:"delivery_geo,omitempty" xml:"delivery_geo,omitempty"`
	// 配送地址
	DeliveryAddress string `json:"delivery_address,omitempty" xml:"delivery_address,omitempty"`
	// 收货人电话
	ConsigneePhone string `json:"consignee_phone,omitempty" xml:"consignee_phone,omitempty"`
	// 运费
	DeliveryFee int64 `json:"delivery_fee,omitempty" xml:"delivery_fee,omitempty"`
	// 是否隐私号 0：非隐私号  1：隐私号
	PrivacyPhoneFlag int64 `json:"privacy_phone_flag,omitempty" xml:"privacy_phone_flag,omitempty"`
}

var poolOrderDeliverer = sync.Pool{
	New: func() any {
		return new(OrderDeliverer)
	},
}

// GetOrderDeliverer() 从对象池中获取OrderDeliverer
func GetOrderDeliverer() *OrderDeliverer {
	return poolOrderDeliverer.Get().(*OrderDeliverer)
}

// ReleaseOrderDeliverer 释放OrderDeliverer
func ReleaseOrderDeliverer(v *OrderDeliverer) {
	v.DelivererCode = ""
	v.DelivererPhone = ""
	v.DelivererName = ""
	v.PickupEndTime = ""
	v.PickupStartTime = ""
	v.BatchEndTime = ""
	v.BatchStartTime = ""
	v.SignTime = ""
	v.DispatchEndTime = ""
	v.DispatchStartTime = ""
	v.PackageEndTime = ""
	v.PackageStartTime = ""
	v.SignMemo = ""
	v.DeliveryStartTime = ""
	v.ConsigneeName = ""
	v.DeliveryEndTime = ""
	v.DeliveryGeo = ""
	v.DeliveryAddress = ""
	v.ConsigneePhone = ""
	v.DeliveryFee = 0
	v.PrivacyPhoneFlag = 0
	poolOrderDeliverer.Put(v)
}
