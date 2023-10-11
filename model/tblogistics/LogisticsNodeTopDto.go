package tblogistics

// LogisticsNodeTopDto 结构体
type LogisticsNodeTopDto struct {
	// ACCEPT(已揽收),TRANSPORT(运输中),DELIVERING(派送中),SIGN(已签收),CANCEL(已取消),FAILED(物流异常)
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 操作时间戳，精确到毫秒（ms）
	OperateTime int64 `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 配送员信息
	Delivery *DeliveryTopDto `json:"delivery,omitempty" xml:"delivery,omitempty"`
	// 货物所在的当前位置
	Location *LocationTopDto `json:"location,omitempty" xml:"location,omitempty"`
}
