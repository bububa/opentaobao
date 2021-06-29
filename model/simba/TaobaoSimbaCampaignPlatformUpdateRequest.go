package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的平台设置 API请求
taobao.simba.campaign.platform.update

更新一个推广计划的平台设置
*/
type TaobaoSimbaCampaignPlatformUpdateRequest struct {
    model.Params
    // 推广计划Id
    _campaignId   int64
    // 搜索投放频道代码数组，频道代码必须是直通车搜索类频道列表中的值。1：淘宝站内搜索，8、无线站内搜索；16:无线站外搜索
    _searchChannels   []int64
    // 非搜索投放频道代码数组，频道代码必须是直通车非搜索类频道列表中的值。1、淘宝站内定向；2、站外定向；8、无线站内定向；16、无线站外定向
    _nonsearchChannels   []int64
    // 已经废弃
    _outsideDiscount   int64
    // 已经废弃
    _mobileDiscount   int64
    // 主人昵称
    _nick   string
}

// 初始化TaobaoSimbaCampaignPlatformUpdateRequest对象
func NewTaobaoSimbaCampaignPlatformUpdateRequest() *TaobaoSimbaCampaignPlatformUpdateRequest{
    return &TaobaoSimbaCampaignPlatformUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.platform.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetCampaignId() int64 {
    return r._campaignId
}
// SearchChannels Setter
// 搜索投放频道代码数组，频道代码必须是直通车搜索类频道列表中的值。1：淘宝站内搜索，8、无线站内搜索；16:无线站外搜索
func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetSearchChannels(_searchChannels []int64) error {
    r._searchChannels = _searchChannels
    r.Set("search_channels", _searchChannels)
    return nil
}

// SearchChannels Getter
func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetSearchChannels() []int64 {
    return r._searchChannels
}
// NonsearchChannels Setter
// 非搜索投放频道代码数组，频道代码必须是直通车非搜索类频道列表中的值。1、淘宝站内定向；2、站外定向；8、无线站内定向；16、无线站外定向
func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetNonsearchChannels(_nonsearchChannels []int64) error {
    r._nonsearchChannels = _nonsearchChannels
    r.Set("nonsearch_channels", _nonsearchChannels)
    return nil
}

// NonsearchChannels Getter
func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetNonsearchChannels() []int64 {
    return r._nonsearchChannels
}
// OutsideDiscount Setter
// 已经废弃
func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetOutsideDiscount(_outsideDiscount int64) error {
    r._outsideDiscount = _outsideDiscount
    r.Set("outside_discount", _outsideDiscount)
    return nil
}

// OutsideDiscount Getter
func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetOutsideDiscount() int64 {
    return r._outsideDiscount
}
// MobileDiscount Setter
// 已经废弃
func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetMobileDiscount(_mobileDiscount int64) error {
    r._mobileDiscount = _mobileDiscount
    r.Set("mobile_discount", _mobileDiscount)
    return nil
}

// MobileDiscount Getter
func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetMobileDiscount() int64 {
    return r._mobileDiscount
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignPlatformUpdateRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignPlatformUpdateRequest) GetNick() string {
    return r._nick
}
