package ott

// LauncherDo 结构体
type LauncherDo struct {
	// 桌面坑位
	Items []Items `json:"items,omitempty" xml:"items>items,omitempty"`
	// 设备属性
	Property *PropertyDo `json:"property,omitempty" xml:"property,omitempty"`
	// 桌面配置
	Version *VersionDo `json:"version,omitempty" xml:"version,omitempty"`
}
