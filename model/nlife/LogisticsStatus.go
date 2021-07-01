package nlife

// LogisticsStatus 结构体
type LogisticsStatus struct {
	// 该物流单里面的商品，商品和商品之间逗号分隔，商品和数量冒号分隔
	GoodsId string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 物流状态:    WAIT_FOR_CONSIGN("有待发货商品"),     WAIT_FOR_SIGN("全部商品已发货"),     SIGNED("全部商品已签收"),     REJECTED("全部商品已拒收");
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 物流公司名称
	LogisticsCompanyName string `json:"logistics_company_name,omitempty" xml:"logistics_company_name,omitempty"`
	// 物流公司ID
	LogisticsCompanyId string `json:"logistics_company_id,omitempty" xml:"logistics_company_id,omitempty"`
	// 物流单号
	LogisticsId string `json:"logistics_id,omitempty" xml:"logistics_id,omitempty"`
	// 更新时间
	ModifyedTime string `json:"modifyed_time,omitempty" xml:"modifyed_time,omitempty"`
	// 更新日志列表
	LogisticsLogList []LogisticsLog `json:"logistics_log_list,omitempty" xml:"logistics_log_list>logistics_log,omitempty"`
}
