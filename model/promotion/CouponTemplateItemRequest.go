package promotion

// CouponTemplateItemRequest 结构体
type CouponTemplateItemRequest struct {
	// 模板表主键id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 券圈品设置
	PromActSkuList []PromActSku `json:"prom_act_sku_list,omitempty" xml:"prom_act_sku_list>prom_act_sku,omitempty"`
	// ump模板ID
	SourceId int64 `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 用户信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
}
