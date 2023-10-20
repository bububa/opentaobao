package yunosappstore

// YunosappstorepadhpapplistResult 结构体
type YunosappstorepadhpapplistResult struct {
	// icon
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// packageName
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// developerName
	DeveloperName string `json:"developer_name,omitempty" xml:"developer_name,omitempty"`
	// uri
	Uri string `json:"uri,omitempty" xml:"uri,omitempty"`
	// url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// intro
	Intro string `json:"intro,omitempty" xml:"intro,omitempty"`
	// versionCode
	VersionCode int64 `json:"version_code,omitempty" xml:"version_code,omitempty"`
}
