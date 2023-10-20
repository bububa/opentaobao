package aesolution

import (
	"sync"
)

// OrderQuery 结构体
type OrderQuery struct {
	// Query order information in multiple order status. For specific order status, see order_status description.
	OrderStatusList []string `json:"order_status_list,omitempty" xml:"order_status_list>string,omitempty"`
	// create date end time.It&#39;s US pacific time
	CreateDateEnd string `json:"create_date_end,omitempty" xml:"create_date_end,omitempty"`
	// create date start time.It&#39;s US pacific time
	CreateDateStart string `json:"create_date_start,omitempty" xml:"create_date_start,omitempty"`
	// modified date start time.It&#39;s US pacific time
	ModifiedDateStart string `json:"modified_date_start,omitempty" xml:"modified_date_start,omitempty"`
	// buyer login id
	BuyerLoginId string `json:"buyer_login_id,omitempty" xml:"buyer_login_id,omitempty"`
	// modified date end time.It&#39;s US pacific time
	ModifiedDateEnd string `json:"modified_date_end,omitempty" xml:"modified_date_end,omitempty"`
	// Order status: PLACE_ORDER_SUCCESS: Waiting for buyer to pay; IN_CANCEL: Buyer request cancellation; WAIT_SELLER_SEND_GOODS: Waiting for your shipment; SELLER_PART_SEND_GOODS: Partial delivery; WAIT_BUYER_ACCEPT_GOODS: Waiting for buyer to receive goods; FUND_PROCESSING: Buyers agree, funds processing; IN_ISSUE : Orders in disputes; IN_FROZEN: Orders in freeze; WAIT_SELLER_EXAMINE_MONEY: Waiting for your confirmation amount; RISK_CONTROL: Orders are in 24 hours of risk control, starting 24 hours after the buyer&#39;s online payment is completed. The above status query can be separately queried separately, and the order status information is not included in the order status. (FINISH, closed order status) FINISH: The completed order needs to be queried separately.
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// Number of pages per page
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// the current page
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

var poolOrderQuery = sync.Pool{
	New: func() any {
		return new(OrderQuery)
	},
}

// GetOrderQuery() 从对象池中获取OrderQuery
func GetOrderQuery() *OrderQuery {
	return poolOrderQuery.Get().(*OrderQuery)
}

// ReleaseOrderQuery 释放OrderQuery
func ReleaseOrderQuery(v *OrderQuery) {
	v.OrderStatusList = v.OrderStatusList[:0]
	v.CreateDateEnd = ""
	v.CreateDateStart = ""
	v.ModifiedDateStart = ""
	v.BuyerLoginId = ""
	v.ModifiedDateEnd = ""
	v.OrderStatus = ""
	v.PageSize = 0
	v.CurrentPage = 0
	poolOrderQuery.Put(v)
}
