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
    nick   string
    // 起始日期
    startDate   string
    // 终止日期
    endDate   string
    // 计划id
    campaignId   int64
    // 推广组id
    adgroupId   int64
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
func (r *TaobaoSubwayAutomatchRptGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSubwayAutomatchRptGetRequest) GetNick() string {
    return r.nick
}
// StartDate Setter
// 起始日期
func (r *TaobaoSubwayAutomatchRptGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoSubwayAutomatchRptGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 终止日期
func (r *TaobaoSubwayAutomatchRptGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoSubwayAutomatchRptGetRequest) GetEndDate() string {
    return r.endDate
}
// CampaignId Setter
// 计划id
func (r *TaobaoSubwayAutomatchRptGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSubwayAutomatchRptGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSubwayAutomatchRptGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSubwayAutomatchRptGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
