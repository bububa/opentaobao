package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广单元小时级别实时报表查询 API请求
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

// 初始化TaobaoSimbaHourReportAdgroupGetRequest对象
func NewTaobaoSimbaHourReportAdgroupGetRequest() *TaobaoSimbaHourReportAdgroupGetRequest{
    return &TaobaoSimbaHourReportAdgroupGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetApiMethodName() string {
    return "taobao.simba.hour.report.adgroup.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetNick() string {
    return r.nick
}
// TheDate Setter
// 时间
func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetTheDate() string {
    return r.theDate
}
// Hour Setter
// 当前小时
func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetHour(hour string) error {
    r.hour = hour
    r.Set("hour", hour)
    return nil
}

// Hour Getter
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetHour() string {
    return r.hour
}
// CampaignId Setter
// 计划id
func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetCampaignId(campaignId string) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetCampaignId() string {
    return r.campaignId
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetAdgroupId(adgroupId string) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetAdgroupId() string {
    return r.adgroupId
}
