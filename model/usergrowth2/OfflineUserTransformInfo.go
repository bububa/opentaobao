package usergrowth2

// OfflineUserTransformInfo 结构体
type OfflineUserTransformInfo struct {
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 新登时间
	TransformTime int64 `json:"transform_time,omitempty" xml:"transform_time,omitempty"`
}
