package product

import (
	"sync"
)

// TaobaoBaikeImportZhubaoPictureResult 结构体
type TaobaoBaikeImportZhubaoPictureResult struct {
	// messageCode
	MessageCode string `json:"message_code,omitempty" xml:"message_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoBaikeImportZhubaoPictureResult = sync.Pool{
	New: func() any {
		return new(TaobaoBaikeImportZhubaoPictureResult)
	},
}

// GetTaobaoBaikeImportZhubaoPictureResult() 从对象池中获取TaobaoBaikeImportZhubaoPictureResult
func GetTaobaoBaikeImportZhubaoPictureResult() *TaobaoBaikeImportZhubaoPictureResult {
	return poolTaobaoBaikeImportZhubaoPictureResult.Get().(*TaobaoBaikeImportZhubaoPictureResult)
}

// ReleaseTaobaoBaikeImportZhubaoPictureResult 释放TaobaoBaikeImportZhubaoPictureResult
func ReleaseTaobaoBaikeImportZhubaoPictureResult(v *TaobaoBaikeImportZhubaoPictureResult) {
	v.MessageCode = ""
	v.Message = ""
	v.Module = ""
	v.Success = false
	poolTaobaoBaikeImportZhubaoPictureResult.Put(v)
}
