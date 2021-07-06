package scbp

// TopP4pQuickEffectQuery 结构体
type TopP4pQuickEffectQuery struct {
	// 结束时间 当inteval=7或30的时候 不需要填写，当inteval=1时需要填写（开始结束时间区间不允许大于180天）
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 开始时间 当inteval=7或30的时候 不需要填写，当inteval=1时需要填写（开始结束时间区间不允许大于180天）
	BeginDate string `json:"begin_date,omitempty" xml:"begin_date,omitempty"`
	// asc(顺序),desc(倒序)
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 排序字段 impression(曝光) click（点击数） ctr（点击率） cost（消耗金额） cpc（平均点击消耗）
	OrderField string `json:"order_field,omitempty" xml:"order_field,omitempty"`
	// 商品名称(可以通过传入商品名称，返回指定商品的数据)
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 填写推广计划名字，单独返回某个计划的数据
	CampaignTitle string `json:"campaign_title,omitempty" xml:"campaign_title,omitempty"`
	// 查询类型，crowd(人群)，region(地域)
	TagType string `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
	// 第几页
	ToPage int64 `json:"to_page,omitempty" xml:"to_page,omitempty"`
	// 每页行数
	PerPageSize int64 `json:"per_page_size,omitempty" xml:"per_page_size,omitempty"`
	// 区间 只能为1 7 30
	Interval int64 `json:"interval,omitempty" xml:"interval,omitempty"`
	// 若查询类型为人群，需要填写此字段  1（潜在访问偏好） 2（潜在采购意向） 3（店铺老客） 4（优选人群）
	TagLabel int64 `json:"tag_label,omitempty" xml:"tag_label,omitempty"`
}
