package ascpffo

// ErrorCode 
type ErrorCode struct {
    // 错误信息
    DisplayMessage   string `json:"display_message,omitempty" xml:"display_message,omitempty"`
    // 标准排查错误码
    Key   string `json:"key,omitempty" xml:"key,omitempty"`
}
