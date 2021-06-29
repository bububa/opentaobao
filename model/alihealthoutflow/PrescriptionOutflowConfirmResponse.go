package alihealthoutflow

// PrescriptionOutflowConfirmResponse 
type PrescriptionOutflowConfirmResponse struct {
    // 备注说明
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // 返回成功或者失败
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
    // 处方url地址
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
}
