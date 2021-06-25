package promotion

// CouponTemplateItemRequest 
type CouponTemplateItemRequest struct {

    // 模板表主键id
    Id   int64 `json:"id,omitempty"`

    // 券圈品设置
    PromActSkuList   []PromActSku `json:"prom_act_sku_list,omitempty"`

    // ump模板ID
    SourceId   int64 `json:"source_id,omitempty"`

    // 用户信息
    UserInfo   *UserInfo `json:"user_info,omitempty"`

}
