package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取城市维度报表 API请求
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

// 初始化TaobaoSimbaReportCityGetRequest对象
func NewTaobaoSimbaReportCityGetRequest() *TaobaoSimbaReportCityGetRequest{
    return &TaobaoSimbaReportCityGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaReportCityGetRequest) GetApiMethodName() string {
    return "taobao.simba.report.city.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaReportCityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaReportCityGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaReportCityGetRequest) GetNick() string {
    return r.nick
}
// TheDate Setter
// 时间
func (r *TaobaoSimbaReportCityGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaReportCityGetRequest) GetTheDate() string {
    return r.theDate
}
// Hour Setter
// 当前小时
func (r *TaobaoSimbaReportCityGetRequest) SetHour(hour string) error {
    r.hour = hour
    r.Set("hour", hour)
    return nil
}

// Hour Getter
func (r TaobaoSimbaReportCityGetRequest) GetHour() string {
    return r.hour
}
// CampaignId Setter
// 计划id
func (r *TaobaoSimbaReportCityGetRequest) SetCampaignId(campaignId string) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaReportCityGetRequest) GetCampaignId() string {
    return r.campaignId
}
