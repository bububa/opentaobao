package game

// AppleTopNewSignNotifyDo 
type AppleTopNewSignNotifyDo struct {

    // 电子卡卡号
    CardNo   string `json:"card_no,omitempty"`

    // TopUpTocken
    Tut   string `json:"tut,omitempty"`

    // 商户订单号
    OrderNo   string `json:"order_no,omitempty"`

    // 附加信息，后续可以扩展
    Memo   string `json:"memo,omitempty"`

    // 商户上送的用户Id
    UserNo   string `json:"user_no,omitempty"`

    // 商户上送UCI
    Uci   string `json:"uci,omitempty"`

    // Mask id
    Mai   string `json:"mai,omitempty"`

    // 网关请求流水号,每次请求唯一
    QueryId   string `json:"query_id,omitempty"`

    // 商户上送tUCI
    Tuci   string `json:"tuci,omitempty"`

}