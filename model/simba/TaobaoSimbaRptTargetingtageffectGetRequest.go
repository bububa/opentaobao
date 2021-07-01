package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取定向效果报表数据 API请求
taobao.simba.rpt.targetingtageffect.get

获取定向效果报表数据
*/
type TaobaoSimbaRptTargetingtageffectGetAPIRequest struct {
    model.Params
    // 被操作者昵称
    _nick   string
    // 计划id
    _campaignId   int64
    // 推广组id
    _adgroupId   int64
    // 起始时间
    _startTime   string
    // 终止时间 ,必须小于今天
    _endTime   string
    // 页面大小
    _pageSize   int64
    // 页码
    _pageNumber   int64
}

// 初始化TaobaoSimbaRptTargetingtageffectGetAPIRequest对象
func NewTaobaoSimbaRptTargetingtageffectGetRequest() *TaobaoSimbaRptTargetingtageffectGetAPIRequest{
    return &TaobaoSimbaRptTargetingtageffectGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.targetingtageffect.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 被操作者昵称
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 计划id
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// StartTime Setter
// 起始时间
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 终止时间 ,必须小于今天
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetEndTime() string {
    return r._endTime
}
// PageSize Setter
// 页面大小
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNumber Setter
// 页码
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetPageNumber(_pageNumber int64) error {
    r._pageNumber = _pageNumber
    r.Set("page_number", _pageNumber)
    return nil
}

// PageNumber Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetPageNumber() int64 {
    return r._pageNumber
}
