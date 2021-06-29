package aliqin

// AlibabaAliqinFcIotQrycardResult 
type AlibabaAliqinFcIotQrycardResult struct {
    // model
    Models   []AlibabaAliqinFcIotQrycardModel `json:"models,omitempty" xml:"models>alibaba_aliqin_fc_iot_qrycard_model,omitempty"`
    // true返回成功，false返回失败
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
}
