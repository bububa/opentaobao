package wdk

import (
	"sync"
)

// OrderPredictQuery 结构体
type OrderPredictQuery struct {
	// 查询日期列表，早于当前时间为查询实际单量，晚于当前时间为预测
	DateList []string `json:"date_list,omitempty" xml:"date_list>string,omitempty"`
	// 配送站code，该参数和warehouse_code二选一
	DeliveryStationCode string `json:"delivery_station_code,omitempty" xml:"delivery_station_code,omitempty"`
	// 门店code，该参数和delivery_station_code二选一
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 数据类型：1）滚动，2）快照
	DataType int64 `json:"data_type,omitempty" xml:"data_type,omitempty"`
	// 订单类型：1）020线上订单，2）b2c常温订单，3）b2c冷链订单
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 时间维度：1）每日一条预测，2）每日48条记录，半小时一条预测
	TimeDimension int64 `json:"time_dimension,omitempty" xml:"time_dimension,omitempty"`
}

var poolOrderPredictQuery = sync.Pool{
	New: func() any {
		return new(OrderPredictQuery)
	},
}

// GetOrderPredictQuery() 从对象池中获取OrderPredictQuery
func GetOrderPredictQuery() *OrderPredictQuery {
	return poolOrderPredictQuery.Get().(*OrderPredictQuery)
}

// ReleaseOrderPredictQuery 释放OrderPredictQuery
func ReleaseOrderPredictQuery(v *OrderPredictQuery) {
	v.DateList = v.DateList[:0]
	v.DeliveryStationCode = ""
	v.WarehouseCode = ""
	v.DataType = 0
	v.OrderType = 0
	v.TimeDimension = 0
	poolOrderPredictQuery.Put(v)
}
