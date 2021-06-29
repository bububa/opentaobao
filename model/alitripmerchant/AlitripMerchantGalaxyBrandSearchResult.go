package alitripmerchant

// AlitripMerchantGalaxyBrandSearchResult 
type AlitripMerchantGalaxyBrandSearchResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 品牌信息
    Brands   *Content `json:"brands,omitempty" xml:"brands,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}