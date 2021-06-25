package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取城市维度报表 APIRequest
taobao.simba.report.city.get

获取城市维度报表
*/
type TaobaoSimbaReportCityGetRequest struct {
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

func NewTaobaoSimbaReportCityGetRequest() *TaobaoSimbaReportCityGetRequest{
    return &TaobaoSimbaReportCityGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaReportCityGetRequest) GetApiMethodName() string {
    return "taobao.simba.report.city.get"
}

func (r TaobaoSimbaReportCityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaReportCityGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaReportCityGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaReportCityGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

func (r TaobaoSimbaReportCityGetRequest) GetTheDate() string {
    return r.theDate
}

func (r *TaobaoSimbaReportCityGetRequest) SetHour(hour string) error {
    r.hour = hour
    r.Set("hour", hour)
    return nil
}

func (r TaobaoSimbaReportCityGetRequest) GetHour() string {
    return r.hour
}

func (r *TaobaoSimbaReportCityGetRequest) SetCampaignId(campaignId string) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaReportCityGetRequest) GetCampaignId() string {
    return r.campaignId
}

