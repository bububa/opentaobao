package wdk

// OrderPredictQuery 
type OrderPredictQuery struct {

    // 数据类型：1）滚动，2）快照
    DataType   int64 `json:"data_type,omitempty"`

    // 查询日期列表，早于当前时间为查询实际单量，晚于当前时间为预测
    DateList   []String `json:"date_list,omitempty"`

    // 配送站code，该参数和warehouse_code二选一
    DeliveryStationCode   string `json:"delivery_station_code,omitempty"`

    // 订单类型：1）020线上订单，2）b2c常温订单，3）b2c冷链订单
    OrderType   int64 `json:"order_type,omitempty"`

    // 时间维度：1）每日一条预测，2）每日48条记录，半小时一条预测
    TimeDimension   int64 `json:"time_dimension,omitempty"`

    // 门店code，该参数和delivery_station_code二选一
    WarehouseCode   string `json:"warehouse_code,omitempty"`

}
