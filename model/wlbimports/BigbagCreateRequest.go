package wlbimports

import (
	"sync"
)

// BigbagCreateRequest 结构体
type BigbagCreateRequest struct {
	// 包裹，快递方案必填
	Packages []HandoverPackageInfo `json:"packages,omitempty" xml:"packages>handover_package_info,omitempty"`
	// 小包LP集合
	ParcelOrderCodeList []string `json:"parcel_order_code_list,omitempty" xml:"parcel_order_code_list>string,omitempty"`
	// 揽件时间范围，快递方案必填
	PickUpTimeRange string `json:"pick_up_time_range,omitempty" xml:"pick_up_time_range,omitempty"`
	// 当日截单时间(18:00)，快递方案必填
	CutOrderTime string `json:"cut_order_time,omitempty" xml:"cut_order_time,omitempty"`
	// 预计送达时间(2021-12-25 12:20:22)，快递方案必填
	EstimatedDeliveryTime string `json:"estimated_delivery_time,omitempty" xml:"estimated_delivery_time,omitempty"`
	// 货物信息（最多70字符），快递必填
	ScItemInfo string `json:"sc_item_info,omitempty" xml:"sc_item_info,omitempty"`
	// pickupType:DOOR_PICKUP 卡车上门揽收 SELF_POST 快递自寄 SELF_SEND 快递自寄
	PickupType string `json:"pickup_type,omitempty" xml:"pickup_type,omitempty"`
	// 保险金额单位，10位以内的正整数
	InsureAmount string `json:"insure_amount,omitempty" xml:"insure_amount,omitempty"`
	// 产品code： EXPRESS:1-2 DAY PREMIUM:NEXT DAY 12:00，快递方案必填
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 揽收资源code
	ReceiveCpCode string `json:"receive_cp_code,omitempty" xml:"receive_cp_code,omitempty"`
	// 期望揽收日期以及时间，快递方案必填
	PlannedShippingDateAndTime string `json:"planned_shipping_date_and_time,omitempty" xml:"planned_shipping_date_and_time,omitempty"`
	// 仓库名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 币种：默认填EUR
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 增值服务类型：SHIPMENT_INSURANCE 保险
	AddServiceType string `json:"add_service_type,omitempty" xml:"add_service_type,omitempty"`
	// 仓code，快递方案必填
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 预约单code，卡车方式必填
	HandoverOrderCode string `json:"handover_order_code,omitempty" xml:"handover_order_code,omitempty"`
	// 收件人，快递方案必填
	ReceiverInfo *ContactInfoRequest `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 发货人信息，快递方案必填
	SenderInfo *ContactInfoRequest `json:"sender_info,omitempty" xml:"sender_info,omitempty"`
	// 是否危险品，快递方案必填
	DangerousFlag bool `json:"dangerous_flag,omitempty" xml:"dangerous_flag,omitempty"`
}

var poolBigbagCreateRequest = sync.Pool{
	New: func() any {
		return new(BigbagCreateRequest)
	},
}

// GetBigbagCreateRequest() 从对象池中获取BigbagCreateRequest
func GetBigbagCreateRequest() *BigbagCreateRequest {
	return poolBigbagCreateRequest.Get().(*BigbagCreateRequest)
}

// ReleaseBigbagCreateRequest 释放BigbagCreateRequest
func ReleaseBigbagCreateRequest(v *BigbagCreateRequest) {
	v.Packages = v.Packages[:0]
	v.ParcelOrderCodeList = v.ParcelOrderCodeList[:0]
	v.PickUpTimeRange = ""
	v.CutOrderTime = ""
	v.EstimatedDeliveryTime = ""
	v.ScItemInfo = ""
	v.PickupType = ""
	v.InsureAmount = ""
	v.ProductCode = ""
	v.ReceiveCpCode = ""
	v.PlannedShippingDateAndTime = ""
	v.StoreName = ""
	v.Currency = ""
	v.AddServiceType = ""
	v.StoreCode = ""
	v.HandoverOrderCode = ""
	v.ReceiverInfo = nil
	v.SellerId = 0
	v.SenderInfo = nil
	v.DangerousFlag = false
	poolBigbagCreateRequest.Put(v)
}
