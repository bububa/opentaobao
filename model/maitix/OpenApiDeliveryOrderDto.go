package maitix

// OpenApiDeliveryOrderDto 结构体
type OpenApiDeliveryOrderDto struct {
	// 快递公司ID
	DeliveryCompanyId int64 `json:"delivery_company_id,omitempty" xml:"delivery_company_id,omitempty"`
	// 快递公司名称
	DeliveryCompanyName string `json:"delivery_company_name,omitempty" xml:"delivery_company_name,omitempty"`
	// 签收时间
	SignTime string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 发货时间
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// 主订单号
	MainOrderId string `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 物流状态，1-发货，2-签收
	DeliveryStatus int64 `json:"delivery_status,omitempty" xml:"delivery_status,omitempty"`
	// 物流单号
	WaybillNo string `json:"waybill_no,omitempty" xml:"waybill_no,omitempty"`
}
