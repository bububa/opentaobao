package ott

// LauncherDO 
type LauncherDO struct {
    // 桌面坑位
    Items   []Items `json:"items,omitempty" xml:"items>items,omitempty"`
    // 设备属性
    Property   *PropertyDO `json:"property,omitempty" xml:"property,omitempty"`
    // 桌面配置
    Version   *VersionDO `json:"version,omitempty" xml:"version,omitempty"`
}
