package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)创建一个推广组 APIRequest
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

func NewTaobaoSimbaSalestarAdgroupAddRequest() *TaobaoSimbaSalestarAdgroupAddRequest{
    return &TaobaoSimbaSalestarAdgroupAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSalestarAdgroupAddRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.adgroup.add"
}

func (r TaobaoSimbaSalestarAdgroupAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSalestarAdgroupAddRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaSalestarAdgroupAddRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaSalestarAdgroupAddRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoSimbaSalestarAdgroupAddRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoSimbaSalestarAdgroupAddRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoSimbaSalestarAdgroupAddRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoSimbaSalestarAdgroupAddRequest) SetImgUrl(imgUrl string) error {
    r.imgUrl = imgUrl
    r.Set("img_url", imgUrl)
    return nil
}

func (r TaobaoSimbaSalestarAdgroupAddRequest) GetImgUrl() string {
    return r.imgUrl
}

