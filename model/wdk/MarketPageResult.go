package wdk

// MarketPageResult 
/* model for simplify = false
type MarketPageResult struct {

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

    // 参加当前活动的商品总数
    
    Total   int64 `json:"total,omitempty"`
    

    // 页面大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 返回的数据
    
    ItemCouponSkuList  struct {
        ItemCouponSku  []ItemCouponSku `json:"item_coupon_sku,omitempty"`
    } `json:"item_coupon_sku_list,omitempty"`
    

    // 当前分页
    
    Current   int64 `json:"current,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // 查询商品是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 查询结果数据
    
    SkuList  struct {
        ItemPoolSku  []ItemPoolSku `json:"item_pool_sku,omitempty"`
    } `json:"sku_list,omitempty"`
    

    // 返回的数据
    
    ItemPoolSkuList  struct {
        ItemPoolSku  []ItemPoolSku `json:"item_pool_sku,omitempty"`
    } `json:"item_pool_sku_list,omitempty"`
    

}
*/

// MarketPageResult 
type MarketPageResult struct {

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 参加当前活动的商品总数
    Total   int64 `json:"total,omitempty"`

    // 页面大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 返回的数据
    ItemCouponSkuList   []ItemCouponSku `json:"item_coupon_sku_list,omitempty"`

    // 当前分页
    Current   int64 `json:"current,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 查询商品是否成功
    Success   bool `json:"success,omitempty"`

    // 查询结果数据
    SkuList   []ItemPoolSku `json:"sku_list,omitempty"`

    // 返回的数据
    ItemPoolSkuList   []ItemPoolSku `json:"item_pool_sku_list,omitempty"`

}
