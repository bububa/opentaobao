package jstsecret

import (
	"sync"
)

// SecretNoBindResponse 结构体
type SecretNoBindResponse struct {
	// 分机号，type=2时有值。拨打时先拨打隐私号再转分机号
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 隐私号码
	SecretNo string `json:"secret_no,omitempty" xml:"secret_no,omitempty"`
	// 隐私号码过期时间
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 隐私号码类型，1=隐私号，2=分机号。
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolSecretNoBindResponse = sync.Pool{
	New: func() any {
		return new(SecretNoBindResponse)
	},
}

// GetSecretNoBindResponse() 从对象池中获取SecretNoBindResponse
func GetSecretNoBindResponse() *SecretNoBindResponse {
	return poolSecretNoBindResponse.Get().(*SecretNoBindResponse)
}

// ReleaseSecretNoBindResponse 释放SecretNoBindResponse
func ReleaseSecretNoBindResponse(v *SecretNoBindResponse) {
	v.Extension = ""
	v.SecretNo = ""
	v.ExpireDate = ""
	v.Type = 0
	poolSecretNoBindResponse.Put(v)
}
