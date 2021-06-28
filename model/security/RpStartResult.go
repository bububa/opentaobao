package security

// RpStartResult 
/* model for simplify = false
type RpStartResult struct {

    // biz
    
    Biz   string `json:"biz,omitempty"`
    

    // extraInfo
    
    ExtraInfo   string `json:"extra_info,omitempty"`
    

    // source
    
    Source   string `json:"source,omitempty"`
    

    // steps
    
    Steps  struct {
        RpStepItem  []RpStepItem `json:"rp_step_item,omitempty"`
    } `json:"steps,omitempty"`
    

    // uploadToken
    
    UploadToken  *struct {
        StsUploadToken  *StsUploadToken `json:"sts_upload_token,omitempty"`
    } `json:"upload_token,omitempty"`
    

}
*/

// RpStartResult 
type RpStartResult struct {

    // biz
    Biz   string `json:"biz,omitempty"`

    // extraInfo
    ExtraInfo   string `json:"extra_info,omitempty"`

    // source
    Source   string `json:"source,omitempty"`

    // steps
    Steps   []RpStepItem `json:"steps,omitempty"`

    // uploadToken
    UploadToken   *StsUploadToken `json:"upload_token,omitempty"`

}
