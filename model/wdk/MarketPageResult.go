package wdk

// MarketPageResult 
type MarketPageResult struct {
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 参加当前活动的商品总数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // 页面大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 返回的数据
    ItemCouponSkuList   []ItemCouponSku `json:"item_coupon_sku_list,omitempty" xml:"item_coupon_sku_list>item_coupon_sku,omitempty"`
    // 当前分页
    Current   int64 `json:"current,omitempty" xml:"current,omitempty"`
    // 总页数
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
    // 查询商品是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 查询结果数据
    SkuList   []ItemPoolSku `json:"sku_list,omitempty" xml:"sku_list>item_pool_sku,omitempty"`
    // 返回的数据
    ItemPoolSkuList   []ItemPoolSku `json:"item_pool_sku_list,omitempty" xml:"item_pool_sku_list>item_pool_sku,omitempty"`
}
