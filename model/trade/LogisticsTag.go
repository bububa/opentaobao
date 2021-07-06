package trade

// LogisticsTag 结构体
type LogisticsTag struct {
	// 服务标签
	LogisticServiceTagList []LogisticServiceTag `json:"logistic_service_tag_list,omitempty" xml:"logistic_service_tag_list>logistic_service_tag,omitempty"`
	// 主订单或子订单的订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
