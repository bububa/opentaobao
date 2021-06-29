package admarket

// AppInfo 
type AppInfo struct {
    // ssp的app名称
    AppName   string `json:"app_name,omitempty" xml:"app_name,omitempty"`
    // ssp的app包名
    Pkg   string `json:"pkg,omitempty" xml:"pkg,omitempty"`
    // app版本号
    VersionCode   int64 `json:"version_code,omitempty" xml:"version_code,omitempty"`
    // app版本名
    VersionName   string `json:"version_name,omitempty" xml:"version_name,omitempty"`
}
