package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索人群实时报表 API请求
taobao.simba.rtrpt.targetingtag.get

获取搜搜人群实时报表
*/
type TaobaoSimbaRtrptTargetingtagGetRequest struct {
    model.Params
    // 旺旺名称
    _nick   string
    // 推广计划id
    _campaignId   int64
    // 推广单元id
    _adgroupId   int64
    // 日期，格式yyyy-mm-dd
    _theDate   string
    // 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
    _trafficType   string
}

// 初始化TaobaoSimbaRtrptTargetingtagGetRequest对象
func NewTaobaoSimbaRtrptTargetingtagGetRequest() *TaobaoSimbaRtrptTargetingtagGetRequest{
    return &TaobaoSimbaRtrptTargetingtagGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.targetingtag.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 旺旺名称
func (r *TaobaoSimbaRtrptTargetingtagGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRtrptTargetingtagGetRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaRtrptTargetingtagGetRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptTargetingtagGetRequest) SetTheDate(_theDate string) error {
    r._theDate = _theDate
    r.Set("the_date", _theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetTheDate() string {
    return r._theDate
}
// TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoSimbaRtrptTargetingtagGetRequest) SetTrafficType(_trafficType string) error {
    r._trafficType = _trafficType
    r.Set("traffic_type", _trafficType)
    return nil
}

// TrafficType Getter
func (r TaobaoSimbaRtrptTargetingtagGetRequest) GetTrafficType() string {
    return r._trafficType
}
