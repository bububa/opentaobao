package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索人群实时报表 APIRequest
taobao.simba.rtrpt.targetingtag.get

获取搜搜人群实时报表
*/
type TaobaoSimbaRtrptTargetingtagGetRequest struct {
    model.Params

    // 旺旺名称
    nick   string 

    // 推广计划id
    campaignId   int64 

    // 推广单元id
    adgroupId   int64 

    // 日期，格式yyyy-mm-dd
    theDate   string 

    // 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
    trafficType   string 

}

func NewTaobaoSimbaRtrptTargetingtagGetRequest() *TaobaoSimbaRtrptTargetingtagGetRequest{
    return &TaobaoSimbaRtrptTargetingtagGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.targetingtag.get"
}

func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRtrptTargetingtagGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRtrptTargetingtagGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaRtrptTargetingtagGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaRtrptTargetingtagGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetTheDate() string {
    return r.theDate
}

func (r *TaobaoSimbaRtrptTargetingtagGetRequest) SetTrafficType(trafficType string) error {
    r.trafficType = trafficType
    r.Set("traffic_type", trafficType)
    return nil
}

func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetTrafficType() string {
    return r.trafficType
}

