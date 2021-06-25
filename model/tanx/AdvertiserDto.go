package tanx

// AdvertiserDto 
type AdvertiserDto struct {

    // 英文名称
    EnglishName   string `json:"english_name,omitempty"`

    // 昵称
    NickName   string `json:"nick_name,omitempty"`

    // 广告主类别（0-淘宝，1-天猫，2-dsp公司，3-dsp个人）
    UserType   int64 `json:"user_type,omitempty"`

    // 广告主名称
    AdvertiserName   string `json:"advertiser_name,omitempty"`

    // 用二进制存储广告主属性1.品牌广告主 2. VIP  4. 世界500强
    AdvertiserType   int64 `json:"advertiser_type,omitempty"`

    // 广告主id
    AdvertiserId   int64 `json:"advertiser_id,omitempty"`

}
