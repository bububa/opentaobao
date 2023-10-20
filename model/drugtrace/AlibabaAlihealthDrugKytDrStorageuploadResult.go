package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrStorageuploadResult 结构体
type AlibabaAlihealthDrugKytDrStorageuploadResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrStorageuploadResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrStorageuploadResult)
	},
}

// GetAlibabaAlihealthDrugKytDrStorageuploadResult() 从对象池中获取AlibabaAlihealthDrugKytDrStorageuploadResult
func GetAlibabaAlihealthDrugKytDrStorageuploadResult() *AlibabaAlihealthDrugKytDrStorageuploadResult {
	return poolAlibabaAlihealthDrugKytDrStorageuploadResult.Get().(*AlibabaAlihealthDrugKytDrStorageuploadResult)
}

// ReleaseAlibabaAlihealthDrugKytDrStorageuploadResult 释放AlibabaAlihealthDrugKytDrStorageuploadResult
func ReleaseAlibabaAlihealthDrugKytDrStorageuploadResult(v *AlibabaAlihealthDrugKytDrStorageuploadResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = false
	v.Success = false
	poolAlibabaAlihealthDrugKytDrStorageuploadResult.Put(v)
}
