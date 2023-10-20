package wdk

import (
	"sync"
)

// AlibabaTclsAelophyOrderReceiptQueryApiResult 结构体
type AlibabaTclsAelophyOrderReceiptQueryApiResult struct {
	// 打印商家/顾客联小票数据列表
	OrderList []ReceiptDto `json:"order_list,omitempty" xml:"order_list>receipt_dto,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}

var poolAlibabaTclsAelophyOrderReceiptQueryApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyOrderReceiptQueryApiResult)
	},
}

// GetAlibabaTclsAelophyOrderReceiptQueryApiResult() 从对象池中获取AlibabaTclsAelophyOrderReceiptQueryApiResult
func GetAlibabaTclsAelophyOrderReceiptQueryApiResult() *AlibabaTclsAelophyOrderReceiptQueryApiResult {
	return poolAlibabaTclsAelophyOrderReceiptQueryApiResult.Get().(*AlibabaTclsAelophyOrderReceiptQueryApiResult)
}

// ReleaseAlibabaTclsAelophyOrderReceiptQueryApiResult 释放AlibabaTclsAelophyOrderReceiptQueryApiResult
func ReleaseAlibabaTclsAelophyOrderReceiptQueryApiResult(v *AlibabaTclsAelophyOrderReceiptQueryApiResult) {
	v.OrderList = v.OrderList[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Status = false
	poolAlibabaTclsAelophyOrderReceiptQueryApiResult.Put(v)
}
