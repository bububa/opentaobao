package xhotelitem

// TaobaoXhotelItemSelectionSellerStatExposureResult 
type TaobaoXhotelItemSelectionSellerStatExposureResult struct {
    // 错误码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 返回结果
    Module   *TaobaoXhotelItemSelectionSellerStatExposureModule `json:"module,omitempty" xml:"module,omitempty"`
    // 接口信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
