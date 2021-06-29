package fans

// PushMessageParamDO 
type PushMessageParamDO struct {
    // 活动id
    ActivityId   string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    // 品牌名
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    // 消息类型
    MessageType   string `json:"message_type,omitempty" xml:"message_type,omitempty"`
    // mixnick
    MixNick   string `json:"mix_nick,omitempty" xml:"mix_nick,omitempty"`
}
