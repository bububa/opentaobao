package qt

// ServiceSubscribe 结构体
type ServiceSubscribe struct {
	// 该用户该收费项目下面的所有的订购记录详情
	UsageDetailList []QualityUsageDetail `json:"usage_detail_list,omitempty" xml:"usage_detail_list>quality_usage_detail,omitempty"`
	// 服务收费项code
	ServiceItemCode string `json:"service_item_code,omitempty" xml:"service_item_code,omitempty"`
	// 订购者昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 过期时间
	GmtExpiry string `json:"gmt_expiry,omitempty" xml:"gmt_expiry,omitempty"`
	// 将要被使用的那条质检订单的价格
	FuturePrice string `json:"future_price,omitempty" xml:"future_price,omitempty"`
	// 订购总数
	AllNum int64 `json:"all_num,omitempty" xml:"all_num,omitempty"`
	// 已经使用的数量
	UsedNum int64 `json:"used_num,omitempty" xml:"used_num,omitempty"`
	// 将要被消耗的质检订单ID
	FutureSubId int64 `json:"future_sub_id,omitempty" xml:"future_sub_id,omitempty"`
	// 可用数量
	AvaliableNum int64 `json:"avaliable_num,omitempty" xml:"avaliable_num,omitempty"`
}
