package aliqin

// AlibabaAliqinFcIotQrycardResult 
/* model for simplify = false
type AlibabaAliqinFcIotQrycardResult struct {

    // model
    
    Models  struct {
        AlibabaAliqinFcIotQrycardModel  []AlibabaAliqinFcIotQrycardModel `json:"alibaba_aliqin_fc_iot_qrycard_model,omitempty"`
    } `json:"models,omitempty"`
    

    // true返回成功，false返回失败
    
    Success   bool `json:"success,omitempty"`
    

    // code
    
    Code   string `json:"code,omitempty"`
    

}
*/

// AlibabaAliqinFcIotQrycardResult 
type AlibabaAliqinFcIotQrycardResult struct {

    // model
    Models   []AlibabaAliqinFcIotQrycardModel `json:"models,omitempty"`

    // true返回成功，false返回失败
    Success   bool `json:"success,omitempty"`

    // code
    Code   string `json:"code,omitempty"`

}
