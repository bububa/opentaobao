package promotion

// LotteryDrawResultDto 
type LotteryDrawResultDto struct {

    // resultType
    ResultType   int64 `json:"result_type,omitempty"`

    // resultMsg
    ResultMsg   string `json:"result_msg,omitempty"`

    // shopTitle
    ShopTitle   string `json:"shop_title,omitempty"`

    // shopLogo
    ShopLogo   string `json:"shop_logo,omitempty"`

    // template
    Template   string `json:"template,omitempty"`

    // award
    Award   *LotteryAwardDto `json:"award,omitempty"`

}
