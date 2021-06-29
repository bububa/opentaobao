package waybill

// WaybillProductType 
type WaybillProductType struct {
    // 产品code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 产品名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 物流服务
    ServiceTypes   []WaybillServiceType `json:"service_types,omitempty" xml:"service_types>waybill_service_type,omitempty"`
}
