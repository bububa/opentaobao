package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的平台设置 APIRequest
taobao.simba.campaign.platform.update

更新一个推广计划的平台设置
*/
type TaobaoSimbaCampaignPlatformUpdateRequest struct {
    model.Params

    // 推广计划Id
    campaignId   int64 

    // 搜索投放频道代码数组，频道代码必须是直通车搜索类频道列表中的值。1：淘宝站内搜索，8、无线站内搜索；16:无线站外搜索
    searchChannels   []Number 

    // 非搜索投放频道代码数组，频道代码必须是直通车非搜索类频道列表中的值。1、淘宝站内定向；2、站外定向；8、无线站内定向；16、无线站外定向
    nonsearchChannels   []Number 

    // 已经废弃
    outsideDiscount   int64 

    // 已经废弃
    mobileDiscount   int64 

    // 主人昵称
    nick   string 

}

func NewTaobaoSimbaCampaignPlatformUpdateRequest() *TaobaoSimbaCampaignPlatformUpdateRequest{
    return &TaobaoSimbaCampaignPlatformUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.platform.update"
}

func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetSearchChannels(searchChannels []Number) error {
    r.searchChannels = searchChannels
    r.Set("search_channels", searchChannels)
    return nil
}

func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetSearchChannels() []Number {
    return r.searchChannels
}

func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetNonsearchChannels(nonsearchChannels []Number) error {
    r.nonsearchChannels = nonsearchChannels
    r.Set("nonsearch_channels", nonsearchChannels)
    return nil
}

func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetNonsearchChannels() []Number {
    return r.nonsearchChannels
}

func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetOutsideDiscount(outsideDiscount int64) error {
    r.outsideDiscount = outsideDiscount
    r.Set("outside_discount", outsideDiscount)
    return nil
}

func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetOutsideDiscount() int64 {
    return r.outsideDiscount
}

func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetMobileDiscount(mobileDiscount int64) error {
    r.mobileDiscount = mobileDiscount
    r.Set("mobile_discount", mobileDiscount)
    return nil
}

func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetMobileDiscount() int64 {
    return r.mobileDiscount
}

func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetNick() string {
    return r.nick
}

