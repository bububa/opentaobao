package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的投放平台设置 API请求
taobao.simba.campaign.platform.get

获得一个推广计划的投放平台设置
*/
type TaobaoSimbaCampaignPlatformGetRequest struct {
    model.Params
    // 主人昵称
    nick   string
    // 推广计划Id
    campaignId   int64
}

// 初始化TaobaoSimbaCampaignPlatformGetRequest对象
func NewTaobaoSimbaCampaignPlatformGetRequest() *TaobaoSimbaCampaignPlatformGetRequest{
    return &TaobaoSimbaCampaignPlatformGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignPlatformGetRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.platform.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignPlatformGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignPlatformGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignPlatformGetRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignPlatformGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaCampaignPlatformGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
