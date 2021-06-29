package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向基础报表 API请求
taobao.simba.rpt.targetingtagbase.get

获取定向基础报表
*/
type TaobaoSimbaRptTargetingtagbaseGetRequest struct {
    model.Params
    // 被操作者昵称
    nick   string
    // 计划id
    campaignId   int64
    // 推广组id
    adgroupId   int64
    // 起始时间
    startTime   string
    // 结束时间
    endTime   string
    // 分页大小
    pageSize   int64
    // 页码
    pageNumber   int64
}

// 初始化TaobaoSimbaRptTargetingtagbaseGetRequest对象
func NewTaobaoSimbaRptTargetingtagbaseGetRequest() *TaobaoSimbaRptTargetingtagbaseGetRequest{
    return &TaobaoSimbaRptTargetingtagbaseGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.targetingtagbase.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 被操作者昵称
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 计划id
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// StartTime Setter
// 起始时间
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetEndTime() string {
    return r.endTime
}
// PageSize Setter
// 分页大小
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNumber Setter
// 页码
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

// PageNumber Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetPageNumber() int64 {
    return r.pageNumber
}
