package cainiaohandover

import (
	"sync"
)

// OpenHandoverContentDetailResponse 结构体
type OpenHandoverContentDetailResponse struct {
	// 大包关联的小包列表
	ParcelOrderList []OpenParcelOrderDto `json:"parcel_order_list,omitempty" xml:"parcel_order_list>open_parcel_order_dto,omitempty"`
	// 交接物物流订单编号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 交接物运单号
	TrackingNumber string `json:"tracking_number,omitempty" xml:"tracking_number,omitempty"`
	// 交接物状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 预估重量
	EstimateWeight string `json:"estimate_weight,omitempty" xml:"estimate_weight,omitempty"`
	// 实际重量
	ActualWeight string `json:"actual_weight,omitempty" xml:"actual_weight,omitempty"`
	// 重量单位
	WeightUnit string `json:"weight_unit,omitempty" xml:"weight_unit,omitempty"`
	// 预估费用
	EstimateFee string `json:"estimate_fee,omitempty" xml:"estimate_fee,omitempty"`
	// 实际费用
	ActualFee string `json:"actual_fee,omitempty" xml:"actual_fee,omitempty"`
	// 费用币种
	FeeCurrency string `json:"fee_currency,omitempty" xml:"fee_currency,omitempty"`
	// 费用单位
	FeeUnit string `json:"fee_unit,omitempty" xml:"fee_unit,omitempty"`
	// 交接物状态
	StatusName string `json:"status_name,omitempty" xml:"status_name,omitempty"`
	// 交接物关联的交接单状态code
	HandoverOrderStatus string `json:"handover_order_status,omitempty" xml:"handover_order_status,omitempty"`
	// 交接物关联的交接单状态名称
	HandoverOrderStatusName string `json:"handover_order_status_name,omitempty" xml:"handover_order_status_name,omitempty"`
}

var poolOpenHandoverContentDetailResponse = sync.Pool{
	New: func() any {
		return new(OpenHandoverContentDetailResponse)
	},
}

// GetOpenHandoverContentDetailResponse() 从对象池中获取OpenHandoverContentDetailResponse
func GetOpenHandoverContentDetailResponse() *OpenHandoverContentDetailResponse {
	return poolOpenHandoverContentDetailResponse.Get().(*OpenHandoverContentDetailResponse)
}

// ReleaseOpenHandoverContentDetailResponse 释放OpenHandoverContentDetailResponse
func ReleaseOpenHandoverContentDetailResponse(v *OpenHandoverContentDetailResponse) {
	v.ParcelOrderList = v.ParcelOrderList[:0]
	v.OrderCode = ""
	v.TrackingNumber = ""
	v.Status = ""
	v.EstimateWeight = ""
	v.ActualWeight = ""
	v.WeightUnit = ""
	v.EstimateFee = ""
	v.ActualFee = ""
	v.FeeCurrency = ""
	v.FeeUnit = ""
	v.StatusName = ""
	v.HandoverOrderStatus = ""
	v.HandoverOrderStatusName = ""
	poolOpenHandoverContentDetailResponse.Put(v)
}
