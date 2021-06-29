package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询流量智选天级报告 API请求
taobao.subway.automatch.rpt.get

查询流量智选天级报告
*/
type TaobaoSubwayAutomatchRptGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 起始日期
    _startDate   string
    // 终止日期
    _endDate   string
    // 计划id
    _campaignId   int64
    // 推广组id
    _adgroupId   int64
}

// 初始化TaobaoSubwayAutomatchRptGetRequest对象
func NewTaobaoSubwayAutomatchRptGetRequest() *TaobaoSubwayAutomatchRptGetRequest{
    return &TaobaoSubwayAutomatchRptGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSubwayAutomatchRptGetRequest) GetApiMethodName() string {
    return "taobao.subway.automatch.rpt.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSubwayAutomatchRptGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSubwayAutomatchRptGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSubwayAutomatchRptGetRequest) GetNick() string {
    return r._nick
}
// StartDate Setter
// 起始日期
func (r *TaobaoSubwayAutomatchRptGetRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSubwayAutomatchRptGetRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 终止日期
func (r *TaobaoSubwayAutomatchRptGetRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSubwayAutomatchRptGetRequest) GetEndDate() string {
    return r._endDate
}
// CampaignId Setter
// 计划id
func (r *TaobaoSubwayAutomatchRptGetRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSubwayAutomatchRptGetRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSubwayAutomatchRptGetRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSubwayAutomatchRptGetRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
