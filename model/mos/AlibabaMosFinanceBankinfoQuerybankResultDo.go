package mos

// AlibabaMosFinanceBankinfoQuerybankResultDo 
/* model for simplify = false
type AlibabaMosFinanceBankinfoQuerybankResultDo struct {

    // 总量
    
    Total   int64 `json:"total,omitempty"`
    

    // 扩展
    
    Extra   string `json:"extra,omitempty"`
    

    // 全链路追踪id
    
    TraceId   string `json:"trace_id,omitempty"`
    

    // 描述
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 是否成功
    
    Data  *struct {
        SupplierBankInfoDto  *SupplierBankInfoDto `json:"supplier_bank_info_dto,omitempty"`
    } `json:"data,omitempty"`
    

    // 成功
    
    ErrCode   int64 `json:"err_code,omitempty"`
    

    // 返回值
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 调用是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaMosFinanceBankinfoQuerybankResultDo 
type AlibabaMosFinanceBankinfoQuerybankResultDo struct {

    // 总量
    Total   int64 `json:"total,omitempty"`

    // 扩展
    Extra   string `json:"extra,omitempty"`

    // 全链路追踪id
    TraceId   string `json:"trace_id,omitempty"`

    // 描述
    ErrMsg   string `json:"err_msg,omitempty"`

    // 是否成功
    Data   *SupplierBankInfoDto `json:"data,omitempty"`

    // 成功
    ErrCode   int64 `json:"err_code,omitempty"`

    // 返回值
    ResultCode   string `json:"result_code,omitempty"`

    // 调用是否成功
    Success   bool `json:"success,omitempty"`

}
