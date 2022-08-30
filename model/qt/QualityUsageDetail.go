package qt

// QualityUsageDetail 结构体
type QualityUsageDetail struct {
	// 质检服务的收费项目码
	ArticleItemCode string `json:"article_item_code,omitempty" xml:"article_item_code,omitempty"`
	// 该订单中每个质检服务的价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 该订单的开通时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 该订单的到期时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 订购ID
	SubId int64 `json:"sub_id,omitempty" xml:"sub_id,omitempty"`
	// 用户ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 该订单订购的数量
	PurchasAmount int64 `json:"purchas_amount,omitempty" xml:"purchas_amount,omitempty"`
	// 该订单已经使用的数量
	ApplicationAmount int64 `json:"application_amount,omitempty" xml:"application_amount,omitempty"`
	// 该订单中尚未使用的数量
	AvailableAmount int64 `json:"available_amount,omitempty" xml:"available_amount,omitempty"`
}
