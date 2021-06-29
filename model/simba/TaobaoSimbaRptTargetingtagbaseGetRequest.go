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
    _nick   string
    // 计划id
    _campaignId   int64
    // 推广组id
    _adgroupId   int64
    // 起始时间
    _startTime   string
    // 结束时间
    _endTime   string
    // 分页大小
    _pageSize   int64
    // 页码
    _pageNumber   int64
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
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 计划id
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// StartTime Setter
// 起始时间
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetEndTime() string {
    return r._endTime
}
// PageSize Setter
// 分页大小
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNumber Setter
// 页码
func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetPageNumber(_pageNumber int64) error {
    r._pageNumber = _pageNumber
    r.Set("page_number", _pageNumber)
    return nil
}

// PageNumber Getter
func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetPageNumber() int64 {
    return r._pageNumber
}
