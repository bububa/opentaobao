package beehive

// TaobaoBeehiveItemCpsUrlResultDo 
type TaobaoBeehiveItemCpsUrlResultDo struct {

    // 商品id和对应的url map
    Model   string `json:"model,omitempty"`

    // 错误码
    Code   string `json:"code,omitempty"`

    // 是否调用成功
    Success   bool `json:"success,omitempty"`

    // 错误信息
    Msg   string `json:"msg,omitempty"`

}
