package mos

// JsonResponse 
type JsonResponse struct {
    // 返回dto
    Data   *SupplierBasisInfoDto `json:"data,omitempty" xml:"data,omitempty"`
    // 报错code
    ErrCode   int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 报错信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // ts
    Ts   int64 `json:"ts,omitempty" xml:"ts,omitempty"`
}
