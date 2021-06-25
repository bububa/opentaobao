package wdk

// ActivitySkuQuery 
type ActivitySkuQuery struct {

    // 五道口活动id
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 分页参数
    Page   *BasePageQuery `json:"page,omitempty"`

    // 商家活动id
    OutActId   string `json:"out_act_id,omitempty"`

    // 需要查询的商品skuCodes
    SkuCodes   []String `json:"sku_codes,omitempty"`

}
