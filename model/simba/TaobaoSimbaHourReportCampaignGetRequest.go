package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计划维度小时报表获取 API请求
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

// 初始化TaobaoSimbaHourReportCampaignGetRequest对象
func NewTaobaoSimbaHourReportCampaignGetRequest() *TaobaoSimbaHourReportCampaignGetRequest{
    return &TaobaoSimbaHourReportCampaignGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaHourReportCampaignGetRequest) GetApiMethodName() string {
    return "taobao.simba.hour.report.campaign.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaHourReportCampaignGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaHourReportCampaignGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaHourReportCampaignGetRequest) GetNick() string {
    return r.nick
}
// TheDate Setter
// 时间
func (r *TaobaoSimbaHourReportCampaignGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaHourReportCampaignGetRequest) GetTheDate() string {
    return r.theDate
}
// Hour Setter
// 当前小时
func (r *TaobaoSimbaHourReportCampaignGetRequest) SetHour(hour string) error {
    r.hour = hour
    r.Set("hour", hour)
    return nil
}

// Hour Getter
func (r TaobaoSimbaHourReportCampaignGetRequest) GetHour() string {
    return r.hour
}
// CampaignId Setter
// 计划id
func (r *TaobaoSimbaHourReportCampaignGetRequest) SetCampaignId(campaignId string) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaHourReportCampaignGetRequest) GetCampaignId() string {
    return r.campaignId
}
