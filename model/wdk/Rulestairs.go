package wdk

// Rulestairs 
type Rulestairs struct {

    // 是否使用满元条件，不能与满件、N件Y折同时使用。此选项为true时，countAt和countBegin必须为false
    AmountAt   bool `json:"amount_at,omitempty"`

    // 满多少元[单位为分，传入700，代表满7元]，amountAt为true时，必须设置
    Amount   int64 `json:"amount,omitempty"`

    // 可换购的数量
    CanExtraItemNum   int64 `json:"can_extra_item_num,omitempty"`

    // 活动是否上不封顶
    EnableMultiple   bool `json:"enable_multiple,omitempty"`

    // 是否使用满件条件，不能与满元、N件Y折同时使用。此选项为true时，countBegin和amountAt必须为false
    CountAt   bool `json:"count_at,omitempty"`

    // 是否使用N件Y折或N件Y元，不能与满元、满件同时使用。此选项为true时，countAt和amountAt必须为false
    CountBegin   bool `json:"count_begin,omitempty"`

    // 满多少件或者第多少件开始参与活动[传入4，代表满4件打折或者第4件打折]。该值与countBegin和countAt相关，如果countBegin为true，则为N件Y折活动，如果countAt为true，则为满件活动
    Count   int64 `json:"count,omitempty"`

    // 是否使用减钱功能，不能与打折功能同时使用，此选项为true时，discount必须为false
    Decrease   bool `json:"decrease,omitempty"`

    // 减多少元[单位为分，传入700，代表减7元]，decrease为true时，必须设置
    DecreaseMoney   int64 `json:"decrease_money,omitempty"`

    // 是否使用打折功能，不能与减钱功能同时使用，此选项为true时，decrease必须为false
    Discount   bool `json:"discount,omitempty"`

    // 打几折[单位为分，传入900，代表打9折]，discount为true时，必须设置
    DiscountRate   int64 `json:"discount_rate,omitempty"`

    // 是否为换购活动
    IsExchange   bool `json:"is_exchange,omitempty"`

    // 是否一口价
    FixPrice   bool `json:"fix_price,omitempty"`

    // 一口价[单位为分]
    FixPriceAmount   int64 `json:"fix_price_amount,omitempty"`

}
