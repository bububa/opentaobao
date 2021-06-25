package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广单元小时级别实时报表查询 APIRequest
taobao.simba.hour.report.adgroup.get

推广单元小时级别实时报表查询
*/
type TaobaoSimbaHourReportAdgroupGetRequest struct {
    model.Params

    // 昵称
    nick   string 

    // 时间
    theDate   string 

    // 当前小时
    hour   string 

    // 计划id
    campaignId   string 

    // 推广单元id
    adgroupId   string 

}

func NewTaobaoSimbaHourReportAdgroupGetRequest() *TaobaoSimbaHourReportAdgroupGetRequest{
    return &TaobaoSimbaHourReportAdgroupGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaHourReportAdgroupGetRequest) GetApiMethodName() string {
    return "taobao.simba.hour.report.adgroup.get"
}

func (r TaobaoSimbaHourReportAdgroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaHourReportAdgroupGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

func (r TaobaoSimbaHourReportAdgroupGetRequest) GetTheDate() string {
    return r.theDate
}

func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetHour(hour string) error {
    r.hour = hour
    r.Set("hour", hour)
    return nil
}

func (r TaobaoSimbaHourReportAdgroupGetRequest) GetHour() string {
    return r.hour
}

func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetCampaignId(campaignId string) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaHourReportAdgroupGetRequest) GetCampaignId() string {
    return r.campaignId
}

func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetAdgroupId(adgroupId string) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaHourReportAdgroupGetRequest) GetAdgroupId() string {
    return r.adgroupId
}

