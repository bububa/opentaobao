package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的投放地域 API请求
taobao.simba.campaign.area.update

更新一个推广计划的投放地域
*/
type TaobaoSimbaCampaignAreaUpdateRequest struct {
    model.Params
    // 推广计划Id
    campaignId   int64
    // 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID；
    area   string
    // 主人昵称
    nick   string
}

// 初始化TaobaoSimbaCampaignAreaUpdateRequest对象
func NewTaobaoSimbaCampaignAreaUpdateRequest() *TaobaoSimbaCampaignAreaUpdateRequest{
    return &TaobaoSimbaCampaignAreaUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignAreaUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.area.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignAreaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignAreaUpdateRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaCampaignAreaUpdateRequest) GetCampaignId() int64 {
    return r.campaignId
}
// Area Setter
// 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID；
func (r *TaobaoSimbaCampaignAreaUpdateRequest) SetArea(area string) error {
    r.area = area
    r.Set("area", area)
    return nil
}

// Area Getter
func (r TaobaoSimbaCampaignAreaUpdateRequest) GetArea() string {
    return r.area
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignAreaUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignAreaUpdateRequest) GetNick() string {
    return r.nick
}
