package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广词实时报表数据 API请求
taobao.simba.rtrpt.bidword.get

获取推广词报表数据
*/
type TaobaoSimbaRtrptBidwordGetRequest struct {
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

// 初始化TaobaoSimbaRtrptBidwordGetRequest对象
func NewTaobaoSimbaRtrptBidwordGetRequest() *TaobaoSimbaRtrptBidwordGetRequest{
    return &TaobaoSimbaRtrptBidwordGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptBidwordGetRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.bidword.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptBidwordGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 用户名
func (r *TaobaoSimbaRtrptBidwordGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRtrptBidwordGetRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRtrptBidwordGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRtrptBidwordGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRtrptBidwordGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRtrptBidwordGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptBidwordGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaRtrptBidwordGetRequest) GetTheDate() string {
    return r.theDate
}
