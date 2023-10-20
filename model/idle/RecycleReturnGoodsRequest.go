package idle

import (
	"sync"
)

// RecycleReturnGoodsRequest 结构体
type RecycleReturnGoodsRequest struct {
	// 快递公司
	LogisticsCompanyName string `json:"logistics_company_name,omitempty" xml:"logistics_company_name,omitempty"`
	// 运单号
	LogisticsMailNo string `json:"logistics_mail_no,omitempty" xml:"logistics_mail_no,omitempty"`
	// 卖家电话
	MobileNumber string `json:"mobile_number,omitempty" xml:"mobile_number,omitempty"`
	// 订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolRecycleReturnGoodsRequest = sync.Pool{
	New: func() any {
		return new(RecycleReturnGoodsRequest)
	},
}

// GetRecycleReturnGoodsRequest() 从对象池中获取RecycleReturnGoodsRequest
func GetRecycleReturnGoodsRequest() *RecycleReturnGoodsRequest {
	return poolRecycleReturnGoodsRequest.Get().(*RecycleReturnGoodsRequest)
}

// ReleaseRecycleReturnGoodsRequest 释放RecycleReturnGoodsRequest
func ReleaseRecycleReturnGoodsRequest(v *RecycleReturnGoodsRequest) {
	v.LogisticsCompanyName = ""
	v.LogisticsMailNo = ""
	v.MobileNumber = ""
	v.BizOrderId = 0
	poolRecycleReturnGoodsRequest.Put(v)
}
