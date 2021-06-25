package util

// AlibabaCuntaoInteractRequisitionGetResult 
type AlibabaCuntaoInteractRequisitionGetResult struct {

    // model
    Model   *ListDto `json:"model,omitempty"`

    // 异常时返回的code
    MsgCode   string `json:"msg_code,omitempty"`

    // 异常时返回的描述
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
