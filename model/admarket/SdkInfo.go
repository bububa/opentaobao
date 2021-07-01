package admarket

// SdkInfo 结构体
type SdkInfo struct {
	// sdk版本名
	VersionName string `json:"version_name,omitempty" xml:"version_name,omitempty"`
	// sdk版本号
	VersionCode int64 `json:"version_code,omitempty" xml:"version_code,omitempty"`
}
