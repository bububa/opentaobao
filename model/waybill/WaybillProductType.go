package waybill

// WaybillProductType 
/* model for simplify = false
type WaybillProductType struct {

    // 产品code
    
    Code   string `json:"code,omitempty"`
    

    // 产品名称
    
    Name   string `json:"name,omitempty"`
    

    // 物流服务
    
    ServiceTypes  struct {
        WaybillServiceType  []WaybillServiceType `json:"waybill_service_type,omitempty"`
    } `json:"service_types,omitempty"`
    

}
*/

// WaybillProductType 
type WaybillProductType struct {

    // 产品code
    Code   string `json:"code,omitempty"`

    // 产品名称
    Name   string `json:"name,omitempty"`

    // 物流服务
    ServiceTypes   []WaybillServiceType `json:"service_types,omitempty"`

}
