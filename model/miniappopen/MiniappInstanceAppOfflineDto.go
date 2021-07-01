package miniappopen

// MiniappInstanceAppOfflineDto 
type MiniappInstanceAppOfflineDto struct {
    // 端信息
    Client   string `json:"client,omitempty" xml:"client,omitempty"`
    // 错误信息
    FailMsg   string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
