package alitripmerchant

// MemberDto 
type MemberDto struct {
    // 用户基本信息
    MemberBaseInfo   *MemberBaseInfoDto `json:"member_base_info,omitempty" xml:"member_base_info,omitempty"`
    // 手机号
    PhoneNum   string `json:"phone_num,omitempty" xml:"phone_num,omitempty"`
    // 微信头像
    WechatAvatarUrl   string `json:"wechat_avatar_url,omitempty" xml:"wechat_avatar_url,omitempty"`
    // 微信昵称
    WechatNickName   string `json:"wechat_nick_name,omitempty" xml:"wechat_nick_name,omitempty"`
    // 是否是会员
    IsMember   bool `json:"is_member,omitempty" xml:"is_member,omitempty"`
    // 手机前缀
    PhonePri   string `json:"phone_pri,omitempty" xml:"phone_pri,omitempty"`
    // 会员卡信息
    CardBaseInfo   *CardBaseInfoDto `json:"card_base_info,omitempty" xml:"card_base_info,omitempty"`
    // 邮箱
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    // token
    Token   string `json:"token,omitempty" xml:"token,omitempty"`
    // 用户id
    TenantUserId   string `json:"tenant_user_id,omitempty" xml:"tenant_user_id,omitempty"`
    // 返回日期毫秒值
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
}
