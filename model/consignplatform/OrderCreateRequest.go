package consignplatform

import (
	"sync"
)

// OrderCreateRequest 结构体
type OrderCreateRequest struct {
	// 子订单列表
	SubOrderList []SubOrderDto `json:"sub_order_list,omitempty" xml:"sub_order_list>sub_order_dto,omitempty"`
	// 商家备注
	UserMemo string `json:"user_memo,omitempty" xml:"user_memo,omitempty"`
	// 外部订单id，会用来做幂等，需要保证唯一
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 买家留言
	BuyerMemo string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
	// 订单来源
	OrderSource string `json:"order_source,omitempty" xml:"order_source,omitempty"`
	// 收件地址
	ReceiveAddress *AddressDtoForTop `json:"receive_address,omitempty" xml:"receive_address,omitempty"`
	// 收件人
	Receiver *PersonDto `json:"receiver,omitempty" xml:"receiver,omitempty"`
}

var poolOrderCreateRequest = sync.Pool{
	New: func() any {
		return new(OrderCreateRequest)
	},
}

// GetOrderCreateRequest() 从对象池中获取OrderCreateRequest
func GetOrderCreateRequest() *OrderCreateRequest {
	return poolOrderCreateRequest.Get().(*OrderCreateRequest)
}

// ReleaseOrderCreateRequest 释放OrderCreateRequest
func ReleaseOrderCreateRequest(v *OrderCreateRequest) {
	v.SubOrderList = v.SubOrderList[:0]
	v.UserMemo = ""
	v.OuterOrderId = ""
	v.BuyerMemo = ""
	v.OrderSource = ""
	v.ReceiveAddress = nil
	v.Receiver = nil
	poolOrderCreateRequest.Put(v)
}
