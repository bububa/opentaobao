package campus

// ControllerOfflineRequestDto 结构体
type ControllerOfflineRequestDto struct {
	// sn
	SnNo string `json:"sn_no,omitempty" xml:"sn_no,omitempty"`
	// 离线日志JSON
	OfflineLog string `json:"offline_log,omitempty" xml:"offline_log,omitempty"`
}
