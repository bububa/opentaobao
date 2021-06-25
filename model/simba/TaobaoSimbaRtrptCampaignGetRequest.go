package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广计划实时报表数据 APIRequest
taobao.simba.rtrpt.campaign.get

获取推广计划实时报表数据
*/
type TaobaoSimbaRtrptCampaignGetRequest struct {
    model.Params

    // 用户名
    nick   string 

    // 日期，格式yyyy-mm-dd
    theDate   string 

}

func NewTaobaoSimbaRtrptCampaignGetRequest() *TaobaoSimbaRtrptCampaignGetRequest{
    return &TaobaoSimbaRtrptCampaignGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRtrptCampaignGetRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.campaign.get"
}

func (r TaobaoSimbaRtrptCampaignGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRtrptCampaignGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRtrptCampaignGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRtrptCampaignGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

func (r TaobaoSimbaRtrptCampaignGetRequest) GetTheDate() string {
    return r.theDate
}

