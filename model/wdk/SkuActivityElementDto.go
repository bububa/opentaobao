package wdk

// SkuActivityElementDto 
type SkuActivityElementDto struct {
    // 商品编码
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    // 限购配置信息
    Limit   *LimitDto `json:"limit,omitempty" xml:"limit,omitempty"`
    // 赠品编码，不填默认赠品为商品本身
    GiftSkuCode   string `json:"gift_sku_code,omitempty" xml:"gift_sku_code,omitempty"`
    // 买N赠M的M参数，赠多少件赠品（目前仅支持买N赠1，giftNum默认为1）
    GiftNum   int64 `json:"gift_num,omitempty" xml:"gift_num,omitempty"`
    // 买N赠M的N参数，买多少件可赠
    BuyNum   int64 `json:"buy_num,omitempty" xml:"buy_num,omitempty"`
    // 一口价金额【分】
    FixPriceMoney   int64 `json:"fix_price_money,omitempty" xml:"fix_price_money,omitempty"`
    // 打折金额【1000位底数】，900代表9折
    DiscountRate   int64 `json:"discount_rate,omitempty" xml:"discount_rate,omitempty"`
    // 减钱金额【分】
    DecreaseMoney   int64 `json:"decrease_money,omitempty" xml:"decrease_money,omitempty"`
}
