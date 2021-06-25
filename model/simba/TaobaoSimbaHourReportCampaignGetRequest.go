package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计划维度小时报表获取 APIRequest
taobao.simba.hour.report.campaign.get

计划维度小时报表获取
*/
type TaobaoSimbaHourReportCampaignGetRequest struct {
    model.Params

    // 昵称
    nick   string 

    // 时间
    theDate   string 

    // 当前小时
    hour   string 

    // 计划id
    campaignId   string 

}

func NewTaobaoSimbaHourReportCampaignGetRequest() *TaobaoSimbaHourReportCampaignGetRequest{
    return &TaobaoSimbaHourReportCampaignGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaHourReportCampaignGetRequest) GetApiMethodName() string {
    return "taobao.simba.hour.report.campaign.get"
}

func (r TaobaoSimbaHourReportCampaignGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaHourReportCampaignGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaHourReportCampaignGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaHourReportCampaignGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

func (r TaobaoSimbaHourReportCampaignGetRequest) GetTheDate() string {
    return r.theDate
}

func (r *TaobaoSimbaHourReportCampaignGetRequest) SetHour(hour string) error {
    r.hour = hour
    r.Set("hour", hour)
    return nil
}

func (r TaobaoSimbaHourReportCampaignGetRequest) GetHour() string {
    return r.hour
}

func (r *TaobaoSimbaHourReportCampaignGetRequest) SetCampaignId(campaignId string) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaHourReportCampaignGetRequest) GetCampaignId() string {
    return r.campaignId
}

