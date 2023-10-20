package trade

import (
	"sync"
)

// TmallAscpOrdersSaleCreateResultDo 结构体
type TmallAscpOrdersSaleCreateResultDo struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// module
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallAscpOrdersSaleCreateResultDo = sync.Pool{
	New: func() any {
		return new(TmallAscpOrdersSaleCreateResultDo)
	},
}

// GetTmallAscpOrdersSaleCreateResultDo() 从对象池中获取TmallAscpOrdersSaleCreateResultDo
func GetTmallAscpOrdersSaleCreateResultDo() *TmallAscpOrdersSaleCreateResultDo {
	return poolTmallAscpOrdersSaleCreateResultDo.Get().(*TmallAscpOrdersSaleCreateResultDo)
}

// ReleaseTmallAscpOrdersSaleCreateResultDo 释放TmallAscpOrdersSaleCreateResultDo
func ReleaseTmallAscpOrdersSaleCreateResultDo(v *TmallAscpOrdersSaleCreateResultDo) {
	v.ErrorMessage = ""
	v.Module = ""
	v.ErrorCode = ""
	v.TotalCount = 0
	v.Success = false
	poolTmallAscpOrdersSaleCreateResultDo.Put(v)
}
