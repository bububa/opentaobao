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
type TaobaoSimbaRptTargetingtageffectGetRequest struct {
    model.Params
    // 被操作者昵称
    nick   string
    // 计划id
    campaignId   int64
    // 推广组id
    adgroupId   int64
    // 起始时间
    startTime   string
    // 终止时间 ,必须小于今天
    endTime   string
    // 页面大小
    pageSize   int64
    // 页码
    pageNumber   int64
}

// 初始化TaobaoSimbaRptTargetingtageffectGetRequest对象
func NewTaobaoSimbaRptTargetingtageffectGetRequest() *TaobaoSimbaRptTargetingtageffectGetRequest{
    return &TaobaoSimbaRptTargetingtageffectGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptTargetingtageffectGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.targetingtageffect.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptTargetingtageffectGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 被操作者昵称
func (r *TaobaoSimbaRptTargetingtageffectGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptTargetingtageffectGetRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 计划id
func (r *TaobaoSimbaRptTargetingtageffectGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptTargetingtageffectGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptTargetingtageffectGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptTargetingtageffectGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// StartTime Setter
// 起始时间
func (r *TaobaoSimbaRptTargetingtageffectGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptTargetingtageffectGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 终止时间 ,必须小于今天
func (r *TaobaoSimbaRptTargetingtageffectGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptTargetingtageffectGetRequest) GetEndTime() string {
    return r.endTime
}
// PageSize Setter
// 页面大小
func (r *TaobaoSimbaRptTargetingtageffectGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptTargetingtageffectGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNumber Setter
// 页码
func (r *TaobaoSimbaRptTargetingtageffectGetRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

// PageNumber Getter
func (r TaobaoSimbaRptTargetingtageffectGetRequest) GetPageNumber() int64 {
    return r.pageNumber
}
