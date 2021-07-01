package security

// OfficialAppVerifyRequest 结构体
type OfficialAppVerifyRequest struct {
	// 应用下载地址
	AppUrl string `json:"app_url,omitempty" xml:"app_url,omitempty"`
	// 应用安装包的MD5或SHA1
	AppHash string `json:"app_hash,omitempty" xml:"app_hash,omitempty"`
	// 应用证书指纹的MD5
	CertMd5 string `json:"cert_md5,omitempty" xml:"cert_md5,omitempty"`
	// 应用包名
	PgkName string `json:"pgk_name,omitempty" xml:"pgk_name,omitempty"`
}
