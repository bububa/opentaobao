package auction

// LatestBids 
type LatestBids struct {

    // 竞买号
    AliasName   string `json:"alias_name,omitempty"`

    // 出价金额(分为单位)
    BidPrice   int64 `json:"bid_price,omitempty"`

    // 出价时间
    BidTime   int64 `json:"bid_time,omitempty"`

    // 标的物类目
    CatName   string `json:"cat_name,omitempty"`

    // 评估价/市场价(分为单位)
    ConsultPrice   int64 `json:"consult_price,omitempty"`

    // 法院名称
    CourtName   string `json:"court_name,omitempty"`

    // 起拍价(分为单位)
    InitPrice   int64 `json:"init_price,omitempty"`

    // 拍品id
    ItemId   int64 `json:"item_id,omitempty"`

    // 拍品标题
    ItemTitle   string `json:"item_title,omitempty"`

    // 拍品图片
    PicUrl   string `json:"pic_url,omitempty"`

}
