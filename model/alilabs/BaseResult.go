package alilabs

// BaseResult 结构体
type BaseResult struct {
	// ret_msg
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// ret_code
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// ret_value
	RetValue *HotWordsContent `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}
