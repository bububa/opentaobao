package waybill

import (
	"sync"
)

// KeyResult 结构体
type KeyResult struct {
	// keyName
	KeyName string `json:"key_name,omitempty" xml:"key_name,omitempty"`
}

var poolKeyResult = sync.Pool{
	New: func() any {
		return new(KeyResult)
	},
}

// GetKeyResult() 从对象池中获取KeyResult
func GetKeyResult() *KeyResult {
	return poolKeyResult.Get().(*KeyResult)
}

// ReleaseKeyResult 释放KeyResult
func ReleaseKeyResult(v *KeyResult) {
	v.KeyName = ""
	poolKeyResult.Put(v)
}
