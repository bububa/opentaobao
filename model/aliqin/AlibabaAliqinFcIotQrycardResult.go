package aliqin

// AlibabaAliqinFcIotQrycardResult 
type AlibabaAliqinFcIotQrycardResult struct {

    // model
    Models   []AlibabaAliqinFcIotQrycardModel `json:"models,omitempty"`

    // true返回成功，false返回失败
    Success   bool `json:"success,omitempty"`

    // code
    Code   string `json:"code,omitempty"`

}
