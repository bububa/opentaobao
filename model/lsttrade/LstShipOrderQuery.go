package lsttrade

import (
	"sync"
)

// LstShipOrderQuery 结构体
type LstShipOrderQuery struct {
	// 配送商名称
	DistributorName string `json:"distributor_name,omitempty" xml:"distributor_name,omitempty"`
	// 主发货单更新时间结束值，格式：yyyy-MM-dd HH:mm:ss
	GmtModifiedEnd string `json:"gmt_modified_end,omitempty" xml:"gmt_modified_end,omitempty"`
	// 主发货单更新时间开始值，格式：yyyy-MM-dd HH:mm:ss
	GmtModifiedStart string `json:"gmt_modified_start,omitempty" xml:"gmt_modified_start,omitempty"`
	// 交易订单号
	MainBizOrderId int64 `json:"main_biz_order_id,omitempty" xml:"main_biz_order_id,omitempty"`
	// 主发货单号
	MainShipOrderId int64 `json:"main_ship_order_id,omitempty" xml:"main_ship_order_id,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页最大主订单数，注意：返回的content_list数据按照子订单维度展开
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}

var poolLstShipOrderQuery = sync.Pool{
	New: func() any {
		return new(LstShipOrderQuery)
	},
}

// GetLstShipOrderQuery() 从对象池中获取LstShipOrderQuery
func GetLstShipOrderQuery() *LstShipOrderQuery {
	return poolLstShipOrderQuery.Get().(*LstShipOrderQuery)
}

// ReleaseLstShipOrderQuery 释放LstShipOrderQuery
func ReleaseLstShipOrderQuery(v *LstShipOrderQuery) {
	v.DistributorName = ""
	v.GmtModifiedEnd = ""
	v.GmtModifiedStart = ""
	v.MainBizOrderId = 0
	v.MainShipOrderId = 0
	v.Page = 0
	v.Size = 0
	poolLstShipOrderQuery.Put(v)
}
