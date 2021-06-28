package wdk

// MaochaoOrderInsuranceQueryResult 
/* model for simplify = false
type MaochaoOrderInsuranceQueryResult struct {

    // 返回码说明
    
    ReturnMsg   string `json:"return_msg,omitempty"`
    

    // 返回码
    
    ReturnCode   string `json:"return_code,omitempty"`
    

    // 是否调用成功
    
    Success   bool `json:"success,omitempty"`
    

    // 返回结果
    
    Model  *struct {
        InsuranceOrder  *InsuranceOrder `json:"insurance_order,omitempty"`
    } `json:"model,omitempty"`
    

}
*/

// MaochaoOrderInsuranceQueryResult 
type MaochaoOrderInsuranceQueryResult struct {

    // 返回码说明
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 返回码
    ReturnCode   string `json:"return_code,omitempty"`

    // 是否调用成功
    Success   bool `json:"success,omitempty"`

    // 返回结果
    Model   *InsuranceOrder `json:"model,omitempty"`

}
