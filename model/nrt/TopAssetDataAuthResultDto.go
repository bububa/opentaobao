package nrt

// TopAssetDataAuthResultDto 结构体
type TopAssetDataAuthResultDto struct {
	// 添加失败的手机号及原因
	Failed []AuthFailedMsg `json:"failed,omitempty" xml:"failed>auth_failed_msg,omitempty"`
}
