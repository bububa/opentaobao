package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取创意实时报表数据 API请求
taobao.simba.rtrpt.creative.get

获取创意实时报表数据
*/
type TaobaoSimbaRtrptCreativeGetAPIRequest struct {
    model.Params
    // 用户名
    _nick   string
    // 推广计划id
    _campaignId   int64
    // 推广组id
    _adgroupId   int64
    // 日期，格式yyyy-mm-dd
    _theDate   string
}

// 初始化TaobaoSimbaRtrptCreativeGetAPIRequest对象
func NewTaobaoSimbaRtrptCreativeGetRequest() *TaobaoSimbaRtrptCreativeGetAPIRequest{
    return &TaobaoSimbaRtrptCreativeGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.creative.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 用户名
func (r *TaobaoSimbaRtrptCreativeGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRtrptCreativeGetAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRtrptCreativeGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptCreativeGetAPIRequest) SetTheDate(_theDate string) error {
    r._theDate = _theDate
    r.Set("the_date", _theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaRtrptCreativeGetAPIRequest) GetTheDate() string {
    return r._theDate
}
