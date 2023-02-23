package alihouse

// IdentityInfoDto 结构体
type IdentityInfoDto struct {
	// 经纪人ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 身份信息
	Identity string `json:"identity,omitempty" xml:"identity,omitempty"`
	// 版本号
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
