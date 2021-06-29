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
    _nick   string
    // 时间
    _theDate   string
    // 当前小时
    _hour   string
    // 计划id
    _campaignId   string
    // 推广单元id
    _adgroupId   string
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
func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetNick() string {
    return r._nick
}
// TheDate Setter
// 时间
func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetTheDate(_theDate string) error {
    r._theDate = _theDate
    r.Set("the_date", _theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetTheDate() string {
    return r._theDate
}
// Hour Setter
// 当前小时
func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetHour(_hour string) error {
    r._hour = _hour
    r.Set("hour", _hour)
    return nil
}

// Hour Getter
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetHour() string {
    return r._hour
}
// CampaignId Setter
// 计划id
func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetCampaignId(_campaignId string) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetCampaignId() string {
    return r._campaignId
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaHourReportAdgroupGetRequest) SetAdgroupId(_adgroupId string) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaHourReportAdgroupGetRequest) GetAdgroupId() string {
    return r._adgroupId
}
