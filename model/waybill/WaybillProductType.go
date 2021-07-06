package waybill

// WaybillProductType 结构体
type WaybillProductType struct {
	// 物流服务
	ServiceTypes []WaybillServiceType `json:"service_types,omitempty" xml:"service_types>waybill_service_type,omitempty"`
	// 产品code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 产品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
