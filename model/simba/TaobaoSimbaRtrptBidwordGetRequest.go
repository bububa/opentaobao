package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广词实时报表数据 APIRequest
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

func NewTaobaoSimbaRtrptBidwordGetRequest() *TaobaoSimbaRtrptBidwordGetRequest{
    return &TaobaoSimbaRtrptBidwordGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRtrptBidwordGetRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.bidword.get"
}

func (r TaobaoSimbaRtrptBidwordGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRtrptBidwordGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRtrptBidwordGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRtrptBidwordGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaRtrptBidwordGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaRtrptBidwordGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaRtrptBidwordGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaRtrptBidwordGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

func (r TaobaoSimbaRtrptBidwordGetRequest) GetTheDate() string {
    return r.theDate
}

