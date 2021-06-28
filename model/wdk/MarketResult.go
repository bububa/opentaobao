package wdk

// MarketResult 
/* model for simplify = false
type MarketResult struct {

    // 处理结果
    
    Message   string `json:"message,omitempty"`
    

    // 错误编码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 结果数据
    
    Datas  struct {
        ItemBuyGiftSku  []ItemBuyGiftSku `json:"item_buy_gift_sku,omitempty"`
    } `json:"datas,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

    // 报名活动成功的商品详情
    
    Data  *struct {
        ItemCouponSku  *ItemCouponSku `json:"item_coupon_sku,omitempty"`
    } `json:"data,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 错误编码
    
    ResultCode   string `json:"result_code,omitempty"`
    

}
*/

// MarketResult 
type MarketResult struct {

    // 处理结果
    Message   string `json:"message,omitempty"`

    // 错误编码
    ErrorCode   string `json:"error_code,omitempty"`

    // 结果数据
    Datas   []ItemBuyGiftSku `json:"datas,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // 报名活动成功的商品详情
    Data   *ItemCouponSku `json:"data,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误编码
    ResultCode   string `json:"result_code,omitempty"`

}
