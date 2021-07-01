package qimen

// TaobaoQimenSupplierSynchronizeResponse 
type TaobaoQimenSupplierSynchronizeResponse struct {
    // success|failure
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // 响应码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 响应信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
