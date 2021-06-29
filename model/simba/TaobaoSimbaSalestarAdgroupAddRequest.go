package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)创建一个推广组 API请求
taobao.simba.salestar.adgroup.add

创建一个推广组
*/
type TaobaoSimbaSalestarAdgroupAddRequest struct {
    model.Params
    // 推广计划Id
    campaignId   int64
    // 商品Id
    itemId   int64
    // 创意标题，最多20个汉字
    title   string
    // 创意图片地址，必须是商品的图片之一
    imgUrl   string
}

// 初始化TaobaoSimbaSalestarAdgroupAddRequest对象
func NewTaobaoSimbaSalestarAdgroupAddRequest() *TaobaoSimbaSalestarAdgroupAddRequest{
    return &TaobaoSimbaSalestarAdgroupAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarAdgroupAddRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.adgroup.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarAdgroupAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaSalestarAdgroupAddRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaSalestarAdgroupAddRequest) GetCampaignId() int64 {
    return r.campaignId
}
// ItemId Setter
// 商品Id
func (r *TaobaoSimbaSalestarAdgroupAddRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoSimbaSalestarAdgroupAddRequest) GetItemId() int64 {
    return r.itemId
}
// Title Setter
// 创意标题，最多20个汉字
func (r *TaobaoSimbaSalestarAdgroupAddRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoSimbaSalestarAdgroupAddRequest) GetTitle() string {
    return r.title
}
// ImgUrl Setter
// 创意图片地址，必须是商品的图片之一
func (r *TaobaoSimbaSalestarAdgroupAddRequest) SetImgUrl(imgUrl string) error {
    r.imgUrl = imgUrl
    r.Set("img_url", imgUrl)
    return nil
}

// ImgUrl Getter
func (r TaobaoSimbaSalestarAdgroupAddRequest) GetImgUrl() string {
    return r.imgUrl
}
