package yunosappstore

// AppInfo 
type AppInfo struct {
    // 应用icon
    Icon   string `json:"icon,omitempty" xml:"icon,omitempty"`
    // 应用包名
    PackageName   string `json:"package_name,omitempty" xml:"package_name,omitempty"`
    // 应用名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 应用版本号
    VersionCode   int64 `json:"version_code,omitempty" xml:"version_code,omitempty"`
    // 应用版本名称
    VersionName   string `json:"version_name,omitempty" xml:"version_name,omitempty"`
}
