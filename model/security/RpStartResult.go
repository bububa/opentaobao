package security

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
