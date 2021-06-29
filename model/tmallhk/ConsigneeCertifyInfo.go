package tmallhk

// ConsigneeCertifyInfo 
type ConsigneeCertifyInfo struct {
    // 身份证正面
    Credential1   string `json:"credential1,omitempty" xml:"credential1,omitempty"`
    // 身份证反面
    Credential2   string `json:"credential2,omitempty" xml:"credential2,omitempty"`
    // 有效期截止时间
    OcrExp   string `json:"ocr_exp,omitempty" xml:"ocr_exp,omitempty"`
    // 身份证号
    OcrId   string `json:"ocr_id,omitempty" xml:"ocr_id,omitempty"`
    // 姓名
    OcrName   string `json:"ocr_name,omitempty" xml:"ocr_name,omitempty"`
    // 订单号
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 证件类型
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
}
