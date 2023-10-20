package security

import (
	"sync"
)

// OfficialAppApplyRequest 结构体
type OfficialAppApplyRequest struct {
	// 应用下载地址
	AppUrl string `json:"app_url,omitempty" xml:"app_url,omitempty"`
	// 应用名
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 应用包名
	PkgName string `json:"pkg_name,omitempty" xml:"pkg_name,omitempty"`
	// 应用官网地址
	Website string `json:"website,omitempty" xml:"website,omitempty"`
	// 开发者名称，企业开发者为公司名称
	Developer string `json:"developer,omitempty" xml:"developer,omitempty"`
	// 应用安装包的MD5或SHA1
	AppHash string `json:"app_hash,omitempty" xml:"app_hash,omitempty"`
	// 应用证书指纹的MD5
	CertMd5 string `json:"cert_md5,omitempty" xml:"cert_md5,omitempty"`
	// 官方版本类型 1-官方正式 2-官方开发版本
	OfficialType int64 `json:"official_type,omitempty" xml:"official_type,omitempty"`
}

var poolOfficialAppApplyRequest = sync.Pool{
	New: func() any {
		return new(OfficialAppApplyRequest)
	},
}

// GetOfficialAppApplyRequest() 从对象池中获取OfficialAppApplyRequest
func GetOfficialAppApplyRequest() *OfficialAppApplyRequest {
	return poolOfficialAppApplyRequest.Get().(*OfficialAppApplyRequest)
}

// ReleaseOfficialAppApplyRequest 释放OfficialAppApplyRequest
func ReleaseOfficialAppApplyRequest(v *OfficialAppApplyRequest) {
	v.AppUrl = ""
	v.AppName = ""
	v.PkgName = ""
	v.Website = ""
	v.Developer = ""
	v.AppHash = ""
	v.CertMd5 = ""
	v.OfficialType = 0
	poolOfficialAppApplyRequest.Put(v)
}
