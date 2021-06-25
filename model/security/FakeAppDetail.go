package security

// FakeAppDetail 
type FakeAppDetail struct {

    // 仿冒应用名称
    AppName   string `json:"app_name,omitempty"`

    // 仿冒应用包名
    PackageName   string `json:"package_name,omitempty"`

    // 仿冒应用感染用户数
    InfectedUsers   int64 `json:"infected_users,omitempty"`

    // 仿冒应用下载地址列表(混淆后的URL)
    DownloadUrls   []String `json:"download_urls,omitempty"`

}
