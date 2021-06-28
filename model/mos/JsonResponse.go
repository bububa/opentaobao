package mos

// JsonResponse 
/* model for simplify = false
type JsonResponse struct {

    // 返回dto
    
    Data  *struct {
        SupplierBasisInfoDto  *SupplierBasisInfoDto `json:"supplier_basis_info_dto,omitempty"`
    } `json:"data,omitempty"`
    

    // 报错code
    
    ErrCode   int64 `json:"err_code,omitempty"`
    

    // 报错信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

    // ts
    
    Ts   int64 `json:"ts,omitempty"`
    

}
*/

// JsonResponse 
type JsonResponse struct {

    // 返回dto
    Data   *SupplierBasisInfoDto `json:"data,omitempty"`

    // 报错code
    ErrCode   int64 `json:"err_code,omitempty"`

    // 报错信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // ts
    Ts   int64 `json:"ts,omitempty"`

}
