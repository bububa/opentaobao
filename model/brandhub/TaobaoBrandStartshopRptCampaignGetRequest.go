package brandhub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺推广计划报表数据查询 API请求
taobao.brand.startshop.rpt.campaign.get

获取明星店铺广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStartshopRptCampaignGetRequest struct {
    model.Params
    // 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
    traffictype   string
    // 查询开始时间(最多查询90天数据)
    startdate   string
    // 查询截至时间(最晚查询到昨天)
    enddate   string
    // 转化周期,默认15天,可选 3,7,15
    effect   int64
}

// 初始化TaobaoBrandStartshopRptCampaignGetRequest对象
func NewTaobaoBrandStartshopRptCampaignGetRequest() *TaobaoBrandStartshopRptCampaignGetRequest{
    return &TaobaoBrandStartshopRptCampaignGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBrandStartshopRptCampaignGetRequest) GetApiMethodName() string {
    return "taobao.brand.startshop.rpt.campaign.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBrandStartshopRptCampaignGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Traffictype Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoBrandStartshopRptCampaignGetRequest) SetTraffictype(traffictype string) error {
    r.traffictype = traffictype
    r.Set("traffictype", traffictype)
    return nil
}

// Traffictype Getter
func (r TaobaoBrandStartshopRptCampaignGetRequest) GetTraffictype() string {
    return r.traffictype
}
// Startdate Setter
// 查询开始时间(最多查询90天数据)
func (r *TaobaoBrandStartshopRptCampaignGetRequest) SetStartdate(startdate string) error {
    r.startdate = startdate
    r.Set("startdate", startdate)
    return nil
}

// Startdate Getter
func (r TaobaoBrandStartshopRptCampaignGetRequest) GetStartdate() string {
    return r.startdate
}
// Enddate Setter
// 查询截至时间(最晚查询到昨天)
func (r *TaobaoBrandStartshopRptCampaignGetRequest) SetEnddate(enddate string) error {
    r.enddate = enddate
    r.Set("enddate", enddate)
    return nil
}

// Enddate Getter
func (r TaobaoBrandStartshopRptCampaignGetRequest) GetEnddate() string {
    return r.enddate
}
// Effect Setter
// 转化周期,默认15天,可选 3,7,15
func (r *TaobaoBrandStartshopRptCampaignGetRequest) SetEffect(effect int64) error {
    r.effect = effect
    r.Set("effect", effect)
    return nil
}

// Effect Getter
func (r TaobaoBrandStartshopRptCampaignGetRequest) GetEffect() int64 {
    return r.effect
}
