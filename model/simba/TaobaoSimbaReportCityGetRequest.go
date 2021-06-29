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
    _nick   string
    // 时间
    _theDate   string
    // 当前小时
    _hour   string
    // 计划id
    _campaignId   string
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
func (r *TaobaoSimbaReportCityGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaReportCityGetRequest) GetNick() string {
    return r._nick
}
// TheDate Setter
// 时间
func (r *TaobaoSimbaReportCityGetRequest) SetTheDate(_theDate string) error {
    r._theDate = _theDate
    r.Set("the_date", _theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaReportCityGetRequest) GetTheDate() string {
    return r._theDate
}
// Hour Setter
// 当前小时
func (r *TaobaoSimbaReportCityGetRequest) SetHour(_hour string) error {
    r._hour = _hour
    r.Set("hour", _hour)
    return nil
}

// Hour Getter
func (r TaobaoSimbaReportCityGetRequest) GetHour() string {
    return r._hour
}
// CampaignId Setter
// 计划id
func (r *TaobaoSimbaReportCityGetRequest) SetCampaignId(_campaignId string) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaReportCityGetRequest) GetCampaignId() string {
    return r._campaignId
}
