package maitix

// OpenProjectDetailDto 结构体
type OpenProjectDetailDto struct {
	// 儿童购票说明
	ChildrenNotice string `json:"children_notice,omitempty" xml:"children_notice,omitempty"`
	// 自助取票
	SelfGetTicketNotice string `json:"self_get_ticket_notice,omitempty" xml:"self_get_ticket_notice,omitempty"`
	// 演出开售时间
	ShowStartTime string `json:"show_start_time,omitempty" xml:"show_start_time,omitempty"`
	// 限购说明
	LimitNotice string `json:"limit_notice,omitempty" xml:"limit_notice,omitempty"`
	// 寄存说明
	DepositInfo string `json:"deposit_info,omitempty" xml:"deposit_info,omitempty"`
	// 演出详情
	ShowDetail string `json:"show_detail,omitempty" xml:"show_detail,omitempty"`
	// 实名制购票
	RealNameNotice string `json:"real_name_notice,omitempty" xml:"real_name_notice,omitempty"`
	// 禁止携带物品说明
	ProhibitedItems string `json:"prohibited_items,omitempty" xml:"prohibited_items,omitempty"`
	// 演出海报图
	ShowPic string `json:"show_pic,omitempty" xml:"show_pic,omitempty"`
	// 退换政策
	PolicyOfReturn string `json:"policy_of_return,omitempty" xml:"policy_of_return,omitempty"`
	// 入场说明
	EntranceNotice string `json:"entrance_notice,omitempty" xml:"entrance_notice,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 二级项目分类名称
	SubClassifyName string `json:"sub_classify_name,omitempty" xml:"sub_classify_name,omitempty"`
	// 是否支持选座
	ChoiceSeatNotice string `json:"choice_seat_notice,omitempty" xml:"choice_seat_notice,omitempty"`
	// 电子票入场提示
	EticketNotice string `json:"eticket_notice,omitempty" xml:"eticket_notice,omitempty"`
	// 演出销售结束时间
	ShowEndTime string `json:"show_end_time,omitempty" xml:"show_end_time,omitempty"`
	// 项目分类名称
	ClassifyName string `json:"classify_name,omitempty" xml:"classify_name,omitempty"`
	// 艺人JSON
	Artists string `json:"artists,omitempty" xml:"artists,omitempty"`
	// 品牌JSON
	IpCard string `json:"ip_card,omitempty" xml:"ip_card,omitempty"`
	// 演出时间
	ShowTime string `json:"show_time,omitempty" xml:"show_time,omitempty"`
	// 取票类型1-无纸化,2-快递票,3-自助换票,4-门店自取。1、3属于电子票换票方式,2、4属于纸质票取票方式
	DeliveryTypes string `json:"delivery_types,omitempty" xml:"delivery_types,omitempty"`
	// 发货城市，国标http://www.mca.gov.cn/article/sj/xzqh/2019/2019/201912251506.html
	PostCity string `json:"post_city,omitempty" xml:"post_city,omitempty"`
	// 上门自取地址
	PickupAddressList string `json:"pickup_address_list,omitempty" xml:"pickup_address_list,omitempty"`
	// 演出时间说明
	PerformTimeDetailList string `json:"perform_time_detail_list,omitempty" xml:"perform_time_detail_list,omitempty"`
	// 项目分类编码
	ClassifyCode int64 `json:"classify_code,omitempty" xml:"classify_code,omitempty"`
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 二级项目分类编码
	SubClassifyCode int64 `json:"sub_classify_code,omitempty" xml:"sub_classify_code,omitempty"`
	// 项目单次限购数量
	PurchaseLimitationOnce int64 `json:"purchase_limitation_once,omitempty" xml:"purchase_limitation_once,omitempty"`
	// 大麦商品id
	DamaiItemId int64 `json:"damai_item_id,omitempty" xml:"damai_item_id,omitempty"`
}
