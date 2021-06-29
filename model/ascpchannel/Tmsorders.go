package ascpchannel

// Tmsorders 
type Tmsorders struct {

    // 运单号
    
    TmsOrderCode   string `json:"tms_order_code,omitempty" xml:"tms_order_code,omitempty"`
    

    // 快递公司code.调用 taobao.logistics.companies.get 获取
    
    TmsServiceCode   string `json:"tms_service_code,omitempty" xml:"tms_service_code,omitempty"`
    

    // 快递公司名称
    
    TmsServiceName   string `json:"tms_service_name,omitempty" xml:"tms_service_name,omitempty"`
    

    // 包裹明细列表
    
    TmsItems   []Tmsitems `json:"tms_items,omitempty" xml:"tms_items,omitempty"`
    

}
