package wdk

// Action 
type Action struct {

    // 减钱金额，单位分
    DecreaseMoney   int64 `json:"decrease_money,omitempty"`

    // 折扣，950，表示9.5折
    DiscountRate   int64 `json:"discount_rate,omitempty"`

    // 是否打折
    Discount   bool `json:"discount,omitempty"`

    // 是否一口价
    FixPrice   bool `json:"fix_price,omitempty"`

    // 一口价金额，单位分
    FixPriceMoney   int64 `json:"fix_price_money,omitempty"`

    // 是否减钱
    Decrease   bool `json:"decrease,omitempty"`

    // 一口价类型，0：普通一口价，1：第N件一口价，2：整体一口价
    FixPriceType   int64 `json:"fix_price_type,omitempty"`

}
