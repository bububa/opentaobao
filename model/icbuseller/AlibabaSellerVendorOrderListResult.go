package icbuseller

// AlibabaSellerVendorOrderListResult 
type AlibabaSellerVendorOrderListResult struct {

    // 接口返回
    
    ExecDescription   string `json:"exec_description,omitempty" xml:"exec_description,omitempty"`
    

    // 返回码
    
    ReturnCode   int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 分页对象
    
    PageDto   *PageDto `json:"page_dto,omitempty" xml:"page_dto,omitempty"`
    

    // 返回对象集合
    
    Dtos   []Dto `json:"dtos,omitempty" xml:"dtos,omitempty"`
    

}
