package ascp

// FuturePlanDetail 结构体
type FuturePlanDetail struct {
	// 表达负卖的店铺宝贝列表
	IcItemList []IcItemDto `json:"ic_item_list,omitempty" xml:"ic_item_list>ic_item_dto,omitempty"`
	// 子订单操作id
	OperationDetailOrderId string `json:"operation_detail_order_id,omitempty" xml:"operation_detail_order_id,omitempty"`
	// 仓code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 负卖计划开始时间。时间格式. yyyy-MM-dd HH:mm:ss
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 负卖计划结束时间。时间格式. yyyy-MM-dd HH:mm:ss
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 可以开始履约时间（预计到仓时间）。 时间格式. yyyy-MM-dd HH:mm:ss
	DeliveryStartDate string `json:"delivery_start_date,omitempty" xml:"delivery_start_date,omitempty"`
	// 货品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 负卖数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 负卖表达策略。1.只表达负卖; 2.实物、负卖一起售卖，先表达实物、再表达负卖
	InventoryStrategy int64 `json:"inventory_strategy,omitempty" xml:"inventory_strategy,omitempty"`
	// 时间策略。1 - 绝对时间， 需要填写 deliveryStartDate 字段; 2 - 相对时间， 需要填写  relativeTime     字段
	TimeStrategy int64 `json:"time_strategy,omitempty" xml:"time_strategy,omitempty"`
	// 相对时间天数(单位:天 适用于相对计划)
	RelativeTime int64 `json:"relative_time,omitempty" xml:"relative_time,omitempty"`
}
