package baoxian

// UploadResult 
type UploadResult struct {

    // model
    Model   *InsAttachmentUploadVo `json:"model,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // isSuccess
    IsSuccess   bool `json:"is_success,omitempty"`

}
