package product

import (
	"sync"
)

// TaobaoBaikeImportZhubaoDataResult 结构体
type TaobaoBaikeImportZhubaoDataResult struct {
	// messageCode
	MessageCode string `json:"message_code,omitempty" xml:"message_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoBaikeImportZhubaoDataResult = sync.Pool{
	New: func() any {
		return new(TaobaoBaikeImportZhubaoDataResult)
	},
}

// GetTaobaoBaikeImportZhubaoDataResult() 从对象池中获取TaobaoBaikeImportZhubaoDataResult
func GetTaobaoBaikeImportZhubaoDataResult() *TaobaoBaikeImportZhubaoDataResult {
	return poolTaobaoBaikeImportZhubaoDataResult.Get().(*TaobaoBaikeImportZhubaoDataResult)
}

// ReleaseTaobaoBaikeImportZhubaoDataResult 释放TaobaoBaikeImportZhubaoDataResult
func ReleaseTaobaoBaikeImportZhubaoDataResult(v *TaobaoBaikeImportZhubaoDataResult) {
	v.MessageCode = ""
	v.Message = ""
	v.Module = ""
	v.Success = false
	poolTaobaoBaikeImportZhubaoDataResult.Put(v)
}
