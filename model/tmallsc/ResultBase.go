package tmallsc

// ResultBase 
type ResultBase struct {

    // gmtModified
    GmtModified   string `json:"gmt_modified,omitempty"`

    // gmtCreate
    GmtCreate   string `json:"gmt_create,omitempty"`

    // value
    Value   string `json:"value,omitempty"`

    // errorCode
    ErrorCode   int64 `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
