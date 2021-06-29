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
type TaobaoSimbaRtrptCreativeGetRequest struct {
    model.Params
    // 用户名
    nick   string
    // 推广计划id
    campaignId   int64
    // 推广组id
    adgroupId   int64
    // 日期，格式yyyy-mm-dd
    theDate   string
}

// 初始化TaobaoSimbaRtrptCreativeGetRequest对象
func NewTaobaoSimbaRtrptCreativeGetRequest() *TaobaoSimbaRtrptCreativeGetRequest{
    return &TaobaoSimbaRtrptCreativeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptCreativeGetRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.creative.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptCreativeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 用户名
func (r *TaobaoSimbaRtrptCreativeGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRtrptCreativeGetRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRtrptCreativeGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRtrptCreativeGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRtrptCreativeGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRtrptCreativeGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptCreativeGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaRtrptCreativeGetRequest) GetTheDate() string {
    return r.theDate
}
