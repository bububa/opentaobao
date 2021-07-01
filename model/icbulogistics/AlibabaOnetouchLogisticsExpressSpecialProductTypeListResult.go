package icbulogistics

// AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult 
type AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult struct {
    // 返回结果描述
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 返回结果编码
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 列表对象
    Values   []SpecialProductTypeDto `json:"values,omitempty" xml:"values>special_product_type_dto,omitempty"`
}
