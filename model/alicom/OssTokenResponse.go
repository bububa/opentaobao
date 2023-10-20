package alicom

import (
	"sync"
)

// OssTokenResponse 结构体
type OssTokenResponse struct {
	// 失效时间
	Expiration string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// token
	SecurityToken string `json:"security_token,omitempty" xml:"security_token,omitempty"`
	// accessKeySecret
	AccessKeySecret string `json:"access_key_secret,omitempty" xml:"access_key_secret,omitempty"`
	// accessKeyId
	AccessKeyId string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
	// status
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolOssTokenResponse = sync.Pool{
	New: func() any {
		return new(OssTokenResponse)
	},
}

// GetOssTokenResponse() 从对象池中获取OssTokenResponse
func GetOssTokenResponse() *OssTokenResponse {
	return poolOssTokenResponse.Get().(*OssTokenResponse)
}

// ReleaseOssTokenResponse 释放OssTokenResponse
func ReleaseOssTokenResponse(v *OssTokenResponse) {
	v.Expiration = ""
	v.SecurityToken = ""
	v.AccessKeySecret = ""
	v.AccessKeyId = ""
	v.Status = ""
	poolOssTokenResponse.Put(v)
}
