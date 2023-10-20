package wdk

import (
	"sync"
)

// TxdBillListGetRequest 结构体
type TxdBillListGetRequest struct {
	// 结束账单日
	EndBillDate string `json:"end_bill_date,omitempty" xml:"end_bill_date,omitempty"`
	// 开始账单日
	StartBillDate string `json:"start_bill_date,omitempty" xml:"start_bill_date,omitempty"`
	// 门店编码
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 页面大小，默认20，最大不得超过200
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页码，默认1
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}

var poolTxdBillListGetRequest = sync.Pool{
	New: func() any {
		return new(TxdBillListGetRequest)
	},
}

// GetTxdBillListGetRequest() 从对象池中获取TxdBillListGetRequest
func GetTxdBillListGetRequest() *TxdBillListGetRequest {
	return poolTxdBillListGetRequest.Get().(*TxdBillListGetRequest)
}

// ReleaseTxdBillListGetRequest 释放TxdBillListGetRequest
func ReleaseTxdBillListGetRequest(v *TxdBillListGetRequest) {
	v.EndBillDate = ""
	v.StartBillDate = ""
	v.ShopCode = ""
	v.PageSize = 0
	v.PageIndex = 0
	poolTxdBillListGetRequest.Put(v)
}
