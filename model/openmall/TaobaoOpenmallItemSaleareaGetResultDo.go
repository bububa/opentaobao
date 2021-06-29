package openmall

// TaobaoOpenmallItemSaleareaGetResultDo 
type TaobaoOpenmallItemSaleareaGetResultDo struct {
    // 可售区域结果
    SaleAreaList   []TopSaleAreaVo `json:"sale_area_list,omitempty" xml:"sale_area_list>top_sale_area_vo,omitempty"`
    // 调用是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
