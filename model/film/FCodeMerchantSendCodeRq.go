package film

// FCodeMerchantSendCodeRq 
type FCodeMerchantSendCodeRq struct {

    // 外部业务用户id
    OutUid   string `json:"out_uid,omitempty"`

    // 每个用户发码的数量
    EachNum   int64 `json:"each_num,omitempty"`

    // 外部商户发码的外部业务号
    ExtOrderId   string `json:"ext_order_id,omitempty"`

    // 发码总数
    Number   int64 `json:"number,omitempty"`

    // 外部下单时间
    OrderTime   string `json:"order_time,omitempty"`

    // 用户ID类型 TAOBAO 或者 TAOBAO_NAME
    UserIdType   string `json:"user_id_type,omitempty"`

    // 外部商户标示
    PartnerCode   string `json:"partner_code,omitempty"`

    // 外面用户昵称
    OutUserName   string `json:"out_user_name,omitempty"`

    // 请求属性字段
    Feature   string `json:"feature,omitempty"`

    // 淘宝用户ID列表，用|分割
    UserIdList   string `json:"user_id_list,omitempty"`

    // 发券码商品mixId
    MixId   string `json:"mix_id,omitempty"`

}
