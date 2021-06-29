package brandhub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺账户报表数据查询 API请求
taobao.brand.startshop.rpt.account.get

获取明星店铺广告主账户整体报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStartshopRptAccountGetRequest struct {
    model.Params
    // 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
    trafficType   string
    // 默认15天
    effect   string
    // 开始时间(最多可查询最近90天)
    endDate   string
    // 截至时间(最晚到昨天)
    startDate   string
}

// 初始化TaobaoBrandStartshopRptAccountGetRequest对象
func NewTaobaoBrandStartshopRptAccountGetRequest() *TaobaoBrandStartshopRptAccountGetRequest{
    return &TaobaoBrandStartshopRptAccountGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBrandStartshopRptAccountGetRequest) GetApiMethodName() string {
    return "taobao.brand.startshop.rpt.account.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBrandStartshopRptAccountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoBrandStartshopRptAccountGetRequest) SetTrafficType(trafficType string) error {
    r.trafficType = trafficType
    r.Set("traffic_type", trafficType)
    return nil
}

// TrafficType Getter
func (r TaobaoBrandStartshopRptAccountGetRequest) GetTrafficType() string {
    return r.trafficType
}
// Effect Setter
// 默认15天
func (r *TaobaoBrandStartshopRptAccountGetRequest) SetEffect(effect string) error {
    r.effect = effect
    r.Set("effect", effect)
    return nil
}

// Effect Getter
func (r TaobaoBrandStartshopRptAccountGetRequest) GetEffect() string {
    return r.effect
}
// EndDate Setter
// 开始时间(最多可查询最近90天)
func (r *TaobaoBrandStartshopRptAccountGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoBrandStartshopRptAccountGetRequest) GetEndDate() string {
    return r.endDate
}
// StartDate Setter
// 截至时间(最晚到昨天)
func (r *TaobaoBrandStartshopRptAccountGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoBrandStartshopRptAccountGetRequest) GetStartDate() string {
    return r.startDate
}
