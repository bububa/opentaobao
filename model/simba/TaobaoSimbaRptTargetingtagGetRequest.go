package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索人群离线报表 API请求
taobao.simba.rpt.targetingtag.get

获取搜搜人群实时报表
*/
type TaobaoSimbaRptTargetingtagGetRequest struct {
    model.Params
    // 用户旺旺名称
    nick   string
    // 推广计划id
    campaignId   int64
    // 推广单元id
    adgroupId   int64
    // 开始时间
    startTime   string
    // 结束时间
    endTime   string
    // 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
    trafficType   string
}

// 初始化TaobaoSimbaRptTargetingtagGetRequest对象
func NewTaobaoSimbaRptTargetingtagGetRequest() *TaobaoSimbaRptTargetingtagGetRequest{
    return &TaobaoSimbaRptTargetingtagGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptTargetingtagGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.targetingtag.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptTargetingtagGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 用户旺旺名称
func (r *TaobaoSimbaRptTargetingtagGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptTargetingtagGetRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRptTargetingtagGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptTargetingtagGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaRptTargetingtagGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptTargetingtagGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// StartTime Setter
// 开始时间
func (r *TaobaoSimbaRptTargetingtagGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptTargetingtagGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间
func (r *TaobaoSimbaRptTargetingtagGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptTargetingtagGetRequest) GetEndTime() string {
    return r.endTime
}
// TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoSimbaRptTargetingtagGetRequest) SetTrafficType(trafficType string) error {
    r.trafficType = trafficType
    r.Set("traffic_type", trafficType)
    return nil
}

// TrafficType Getter
func (r TaobaoSimbaRptTargetingtagGetRequest) GetTrafficType() string {
    return r.trafficType
}
