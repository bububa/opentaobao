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
    nick   string
    // 推广计划Id
    campaignId   int64
    // 商品Id
    itemId   int64
    // 推广组默认出价，单位为分，不能小于5 不能大于日最高限额
    defaultPrice   int64
    // 创意标题，最多20个汉字
    title   string
    // 创意图片地址，必须是商品的图片之一
    imgUrl   string
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
func (r *TaobaoSimbaAdgroupAddRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupAddRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaAdgroupAddRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaAdgroupAddRequest) GetCampaignId() int64 {
    return r.campaignId
}
// ItemId Setter
// 商品Id
func (r *TaobaoSimbaAdgroupAddRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoSimbaAdgroupAddRequest) GetItemId() int64 {
    return r.itemId
}
// DefaultPrice Setter
// 推广组默认出价，单位为分，不能小于5 不能大于日最高限额
func (r *TaobaoSimbaAdgroupAddRequest) SetDefaultPrice(defaultPrice int64) error {
    r.defaultPrice = defaultPrice
    r.Set("default_price", defaultPrice)
    return nil
}

// DefaultPrice Getter
func (r TaobaoSimbaAdgroupAddRequest) GetDefaultPrice() int64 {
    return r.defaultPrice
}
// Title Setter
// 创意标题，最多20个汉字
func (r *TaobaoSimbaAdgroupAddRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoSimbaAdgroupAddRequest) GetTitle() string {
    return r.title
}
// ImgUrl Setter
// 创意图片地址，必须是商品的图片之一
func (r *TaobaoSimbaAdgroupAddRequest) SetImgUrl(imgUrl string) error {
    r.imgUrl = imgUrl
    r.Set("img_url", imgUrl)
    return nil
}

// ImgUrl Getter
func (r TaobaoSimbaAdgroupAddRequest) GetImgUrl() string {
    return r.imgUrl
}
