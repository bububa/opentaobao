package dutyfree

import (
	"sync"
)

// AlibabaDutyfreeStockQueryResultDo 结构体
type AlibabaDutyfreeStockQueryResultDo struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 具体库存信息
	Object *StockResultDto `json:"object,omitempty" xml:"object,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaDutyfreeStockQueryResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaDutyfreeStockQueryResultDo)
	},
}

// GetAlibabaDutyfreeStockQueryResultDo() 从对象池中获取AlibabaDutyfreeStockQueryResultDo
func GetAlibabaDutyfreeStockQueryResultDo() *AlibabaDutyfreeStockQueryResultDo {
	return poolAlibabaDutyfreeStockQueryResultDo.Get().(*AlibabaDutyfreeStockQueryResultDo)
}

// ReleaseAlibabaDutyfreeStockQueryResultDo 释放AlibabaDutyfreeStockQueryResultDo
func ReleaseAlibabaDutyfreeStockQueryResultDo(v *AlibabaDutyfreeStockQueryResultDo) {
	v.Message = ""
	v.Code = 0
	v.Object = nil
	v.Success = false
	poolAlibabaDutyfreeStockQueryResultDo.Put(v)
}
