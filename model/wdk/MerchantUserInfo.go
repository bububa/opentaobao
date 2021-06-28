package wdk

// MerchantUserInfo 
/* model for simplify = false
type MerchantUserInfo struct {

    // 真实手机号
    
    RealPhone   string `json:"real_phone,omitempty"`
    

    // 生日
    
    Birthday   string `json:"birthday,omitempty"`
    

    // 地址
    
    Address   string `json:"address,omitempty"`
    

    // 性别
    
    Gender   string `json:"gender,omitempty"`
    

    // 注册时间
    
    RegisterTime   string `json:"register_time,omitempty"`
    

    // 等级
    
    MemberLevel   string `json:"member_level,omitempty"`
    

    // 入会来源
    
    Source   string `json:"source,omitempty"`
    

    // 用户名
    
    UserName   string `json:"user_name,omitempty"`
    

    // 卡号
    
    CardNo   string `json:"card_no,omitempty"`
    

    // 积分余额，可以为小数
    
    ScoreBalance   string `json:"score_balance,omitempty"`
    

    // 自定义渠道类型
    
    CustomizeChannel   string `json:"customize_channel,omitempty"`
    

    // 商家侧统一用户标识ID,如统一会员id
    
    UnionUid   string `json:"union_uid,omitempty"`
    

    // 渠道用户ID，如erp会员id，淘宝openid、饿了么uid、微信openId
    
    ChannelUserId   string `json:"channel_user_id,omitempty"`
    

    // 储值卡余额，需要统一单位为”分“，然后取整上传
    
    CardBalance   int64 `json:"card_balance,omitempty"`
    

    // 扩项属性
    
    ExtendProperty   string `json:"extend_property,omitempty"`
    

    // 会员开始时间
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 会员结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 是否续费
    
    Renew   bool `json:"renew,omitempty"`
    

    // 引导来源标识，填写引导来源标识备注，如门店编码
    
    SourceTag   string `json:"source_tag,omitempty"`
    

    // 是否付费
    
    PayMember   bool `json:"pay_member,omitempty"`
    

    // storeMember：门店会员，txd：淘鲜达，eleme：饿了么，weixin：微信小程序
    
    ChannelCode   string `json:"channel_code,omitempty"`
    

    // isv系统中的完整的用户信息
    
    OriginWholeData   string `json:"origin_whole_data,omitempty"`
    

    // 微信名
    
    WxUserName   string `json:"wx_user_name,omitempty"`
    

    // 微信unionId
    
    WxUnionId   string `json:"wx_union_id,omitempty"`
    

}
*/

// MerchantUserInfo 
type MerchantUserInfo struct {

    // 真实手机号
    RealPhone   string `json:"real_phone,omitempty"`

    // 生日
    Birthday   string `json:"birthday,omitempty"`

    // 地址
    Address   string `json:"address,omitempty"`

    // 性别
    Gender   string `json:"gender,omitempty"`

    // 注册时间
    RegisterTime   string `json:"register_time,omitempty"`

    // 等级
    MemberLevel   string `json:"member_level,omitempty"`

    // 入会来源
    Source   string `json:"source,omitempty"`

    // 用户名
    UserName   string `json:"user_name,omitempty"`

    // 卡号
    CardNo   string `json:"card_no,omitempty"`

    // 积分余额，可以为小数
    ScoreBalance   string `json:"score_balance,omitempty"`

    // 自定义渠道类型
    CustomizeChannel   string `json:"customize_channel,omitempty"`

    // 商家侧统一用户标识ID,如统一会员id
    UnionUid   string `json:"union_uid,omitempty"`

    // 渠道用户ID，如erp会员id，淘宝openid、饿了么uid、微信openId
    ChannelUserId   string `json:"channel_user_id,omitempty"`

    // 储值卡余额，需要统一单位为”分“，然后取整上传
    CardBalance   int64 `json:"card_balance,omitempty"`

    // 扩项属性
    ExtendProperty   string `json:"extend_property,omitempty"`

    // 会员开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 会员结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 是否续费
    Renew   bool `json:"renew,omitempty"`

    // 引导来源标识，填写引导来源标识备注，如门店编码
    SourceTag   string `json:"source_tag,omitempty"`

    // 是否付费
    PayMember   bool `json:"pay_member,omitempty"`

    // storeMember：门店会员，txd：淘鲜达，eleme：饿了么，weixin：微信小程序
    ChannelCode   string `json:"channel_code,omitempty"`

    // isv系统中的完整的用户信息
    OriginWholeData   string `json:"origin_whole_data,omitempty"`

    // 微信名
    WxUserName   string `json:"wx_user_name,omitempty"`

    // 微信unionId
    WxUnionId   string `json:"wx_union_id,omitempty"`

}
