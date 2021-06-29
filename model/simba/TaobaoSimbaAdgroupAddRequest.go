package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建一个推广组 API请求
taobao.simba.adgroup.add

创建一个推广组
*/
type TaobaoSimbaAdgroupAddRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广计划Id
    _campaignId   int64
    // 商品Id
    _itemId   int64
    // 推广组默认出价，单位为分，不能小于5 不能大于日最高限额
    _defaultPrice   int64
    // 创意标题，最多20个汉字
    _title   string
    // 创意图片地址，必须是商品的图片之一
    _imgUrl   string
}

// 初始化TaobaoSimbaAdgroupAddRequest对象
func NewTaobaoSimbaAdgroupAddRequest() *TaobaoSimbaAdgroupAddRequest{
    return &TaobaoSimbaAdgroupAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupAddRequest) GetApiMethodName() string {
    return "taobao.simba.adgroup.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupAddRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupAddRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaAdgroupAddRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaAdgroupAddRequest) GetCampaignId() int64 {
    return r._campaignId
}
// ItemId Setter
// 商品Id
func (r *TaobaoSimbaAdgroupAddRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoSimbaAdgroupAddRequest) GetItemId() int64 {
    return r._itemId
}
// DefaultPrice Setter
// 推广组默认出价，单位为分，不能小于5 不能大于日最高限额
func (r *TaobaoSimbaAdgroupAddRequest) SetDefaultPrice(_defaultPrice int64) error {
    r._defaultPrice = _defaultPrice
    r.Set("default_price", _defaultPrice)
    return nil
}

// DefaultPrice Getter
func (r TaobaoSimbaAdgroupAddRequest) GetDefaultPrice() int64 {
    return r._defaultPrice
}
// Title Setter
// 创意标题，最多20个汉字
func (r *TaobaoSimbaAdgroupAddRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoSimbaAdgroupAddRequest) GetTitle() string {
    return r._title
}
// ImgUrl Setter
// 创意图片地址，必须是商品的图片之一
func (r *TaobaoSimbaAdgroupAddRequest) SetImgUrl(_imgUrl string) error {
    r._imgUrl = _imgUrl
    r.Set("img_url", _imgUrl)
    return nil
}

// ImgUrl Getter
func (r TaobaoSimbaAdgroupAddRequest) GetImgUrl() string {
    return r._imgUrl
}
