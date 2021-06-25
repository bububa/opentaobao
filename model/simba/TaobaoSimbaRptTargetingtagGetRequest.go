package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索人群离线报表 APIRequest
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

func NewTaobaoSimbaRptTargetingtagGetRequest() *TaobaoSimbaRptTargetingtagGetRequest{
    return &TaobaoSimbaRptTargetingtagGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRptTargetingtagGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.targetingtag.get"
}

func (r TaobaoSimbaRptTargetingtagGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRptTargetingtagGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRptTargetingtagGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRptTargetingtagGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaRptTargetingtagGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaRptTargetingtagGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaRptTargetingtagGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaRptTargetingtagGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaRptTargetingtagGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaRptTargetingtagGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSimbaRptTargetingtagGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSimbaRptTargetingtagGetRequest) SetTrafficType(trafficType string) error {
    r.trafficType = trafficType
    r.Set("traffic_type", trafficType)
    return nil
}

func (r TaobaoSimbaRptTargetingtagGetRequest) GetTrafficType() string {
    return r.trafficType
}

