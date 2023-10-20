package security

import (
	"sync"
)

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

var poolOfficialAppVerifyRequest = sync.Pool{
	New: func() any {
		return new(OfficialAppVerifyRequest)
	},
}

// GetOfficialAppVerifyRequest() 从对象池中获取OfficialAppVerifyRequest
func GetOfficialAppVerifyRequest() *OfficialAppVerifyRequest {
	return poolOfficialAppVerifyRequest.Get().(*OfficialAppVerifyRequest)
}

// ReleaseOfficialAppVerifyRequest 释放OfficialAppVerifyRequest
func ReleaseOfficialAppVerifyRequest(v *OfficialAppVerifyRequest) {
	v.AppUrl = ""
	v.AppHash = ""
	v.CertMd5 = ""
	v.PgkName = ""
	poolOfficialAppVerifyRequest.Put(v)
}
