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
type TaobaoSimbaRptTargetingtagbaseGetAPIRequest struct {
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

// 初始化TaobaoSimbaRptTargetingtagbaseGetAPIRequest对象
func NewTaobaoSimbaRptTargetingtagbaseGetRequest() *TaobaoSimbaRptTargetingtagbaseGetAPIRequest{
    return &TaobaoSimbaRptTargetingtagbaseGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.targetingtagbase.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 被操作者昵称
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 计划id
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// StartTime Setter
// 起始时间
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetEndTime() string {
    return r._endTime
}
// PageSize Setter
// 分页大小
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNumber Setter
// 页码
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetPageNumber(_pageNumber int64) error {
    r._pageNumber = _pageNumber
    r.Set("page_number", _pageNumber)
    return nil
}

// PageNumber Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetPageNumber() int64 {
    return r._pageNumber
}
