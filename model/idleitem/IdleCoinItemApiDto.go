package idleitem

// IdleCoinItemApiDto 
type IdleCoinItemApiDto struct {
    // 描述
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // 标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 用户昵称
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
    // 视频Id
    VideoId   string `json:"video_id,omitempty" xml:"video_id,omitempty"`
    // 类目
    CateId   int64 `json:"cate_id,omitempty" xml:"cate_id,omitempty"`
    // 视频封面图ID
    VideoCoverId   int64 `json:"video_cover_id,omitempty" xml:"video_cover_id,omitempty"`
    // 免费送类型
    IdleCoinType   int64 `json:"idle_coin_type,omitempty" xml:"idle_coin_type,omitempty"`
    // 起拍价
    BidReservePrice   int64 `json:"bid_reserve_price,omitempty" xml:"bid_reserve_price,omitempty"`
    // 持续时间
    BidInterval   int64 `json:"bid_interval,omitempty" xml:"bid_interval,omitempty"`
    // 一口价的价格
    BuynowReservePrice   int64 `json:"buynow_reserve_price,omitempty" xml:"buynow_reserve_price,omitempty"`
    // 图片ID
    ImageIds   []int64 `json:"image_ids,omitempty" xml:"image_ids>int64,omitempty"`
    // 库存数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 新旧程度
    StuffStatus   int64 `json:"stuff_status,omitempty" xml:"stuff_status,omitempty"`
    // 地址信息
    AddressDto   *RentAddressDto `json:"address_dto,omitempty" xml:"address_dto,omitempty"`
    // 运费
    PostPrice   int64 `json:"post_price,omitempty" xml:"post_price,omitempty"`
    // 扫描码
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    // 扫描码商品名称
    BarcodeName   string `json:"barcode_name,omitempty" xml:"barcode_name,omitempty"`
    // 扩展信息
    ExtraInfo   *Extrainfo `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
}
