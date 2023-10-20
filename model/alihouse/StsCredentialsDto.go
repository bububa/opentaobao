package alihouse

import (
	"sync"
)

// StsCredentialsDto 结构体
type StsCredentialsDto struct {
	// oss访问临时ak
	AccessKeyId string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
	// oss访问临时安全token
	SecurityToken string `json:"security_token,omitempty" xml:"security_token,omitempty"`
	// oss访问临时sk
	AccessKeySecret string `json:"access_key_secret,omitempty" xml:"access_key_secret,omitempty"`
}

var poolStsCredentialsDto = sync.Pool{
	New: func() any {
		return new(StsCredentialsDto)
	},
}

// GetStsCredentialsDto() 从对象池中获取StsCredentialsDto
func GetStsCredentialsDto() *StsCredentialsDto {
	return poolStsCredentialsDto.Get().(*StsCredentialsDto)
}

// ReleaseStsCredentialsDto 释放StsCredentialsDto
func ReleaseStsCredentialsDto(v *StsCredentialsDto) {
	v.AccessKeyId = ""
	v.SecurityToken = ""
	v.AccessKeySecret = ""
	poolStsCredentialsDto.Put(v)
}
