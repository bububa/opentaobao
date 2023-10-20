package smartstore

// TmallpopupstoreactivitydevicequeryResultDto 结构体
type TmallpopupstoreactivitydevicequeryResultDto struct {
	// 结果code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 结果数据条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 返回结果
	Result *TmallpopupstoreactivitydevicequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
