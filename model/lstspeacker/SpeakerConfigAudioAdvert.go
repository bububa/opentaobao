package lstspeacker

// SpeakerConfigAudioAdvert 
type SpeakerConfigAudioAdvert struct {
    // 广告ID
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    // 订单ID
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 播放时间戳
    PlayScheduleTime   int64 `json:"play_schedule_time,omitempty" xml:"play_schedule_time,omitempty"`
    // 名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // md5
    Md5   string `json:"md5,omitempty" xml:"md5,omitempty"`
    // 文件大小kb
    FileSize   int64 `json:"file_size,omitempty" xml:"file_size,omitempty"`
    // 播放时长
    Length   int64 `json:"length,omitempty" xml:"length,omitempty"`
    // 单价元
    UnitPrice   string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
    // 折扣率
    AdvertDiscountRatio   string `json:"advert_discount_ratio,omitempty" xml:"advert_discount_ratio,omitempty"`
    // 门店折扣率
    StoreDiscountRatio   string `json:"store_discount_ratio,omitempty" xml:"store_discount_ratio,omitempty"`
    // 广告地址
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
}
