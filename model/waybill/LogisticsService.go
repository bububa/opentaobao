package waybill

// LogisticsService 
/* model for simplify = false
type LogisticsService struct {

    // 服务类型值，json格式表示
    
    ServiceValue4Json   string `json:"service_value4_json,omitempty"`
    

    // 服务编码
    
    ServiceCode   string `json:"service_code,omitempty"`
    

}
*/

// LogisticsService 
type LogisticsService struct {

    // 服务类型值，json格式表示
    ServiceValue4Json   string `json:"service_value4_json,omitempty"`

    // 服务编码
    ServiceCode   string `json:"service_code,omitempty"`

}
