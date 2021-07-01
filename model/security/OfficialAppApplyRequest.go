package security

// OfficialAppApplyRequest 结构体
type OfficialAppApplyRequest struct {
	// 应用下载地址
	AppUrl string `json:"app_url,omitempty" xml:"app_url,omitempty"`
	// 应用安装包的MD5或SHA1
	AppHash string `json:"app_hash,omitempty" xml:"app_hash,omitempty"`
	// 官方版本类型 1-官方正式 2-官方开发版本
	OfficialType int64 `json:"official_type,omitempty" xml:"official_type,omitempty"`
	// 应用名
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 应用包名
	PkgName string `json:"pkg_name,omitempty" xml:"pkg_name,omitempty"`
	// 应用官网地址
	Website string `json:"website,omitempty" xml:"website,omitempty"`
	// 开发者名称，企业开发者为公司名称
	Developer string `json:"developer,omitempty" xml:"developer,omitempty"`
	// 应用证书指纹的MD5
	CertMd5 string `json:"cert_md5,omitempty" xml:"cert_md5,omitempty"`
}
