package wdk

// ActivitySkuQuery 
/* model for simplify = false
type ActivitySkuQuery struct {

    // 五道口活动id
    
    ActivityId   int64 `json:"activity_id,omitempty"`
    

    // 分页参数
    
    Page  *struct {
        BasePageQuery  *BasePageQuery `json:"base_page_query,omitempty"`
    } `json:"page,omitempty"`
    

    // 商家活动id
    
    OutActId   string `json:"out_act_id,omitempty"`
    

    // 需要查询的商品skuCodes
    
    SkuCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sku_codes,omitempty"`
    

}
*/

// ActivitySkuQuery 
type ActivitySkuQuery struct {

    // 五道口活动id
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 分页参数
    Page   *BasePageQuery `json:"page,omitempty"`

    // 商家活动id
    OutActId   string `json:"out_act_id,omitempty"`

    // 需要查询的商品skuCodes
    SkuCodes   []string `json:"sku_codes,omitempty"`

}
