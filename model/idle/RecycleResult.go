package idle

// RecycleResult 结构体
type RecycleResult struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// errMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 闲鱼后台spu状态，0为在线，1为待测试，需要测试完成后再次挂载，-1为已删除（置为无效）
	SpuStatus int64 `json:"spu_status,omitempty" xml:"spu_status,omitempty"`
}
