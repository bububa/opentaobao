package feedflow

// LaunchAreaDto 结构体
type LaunchAreaDto struct {
	// 地址名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 地址code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
