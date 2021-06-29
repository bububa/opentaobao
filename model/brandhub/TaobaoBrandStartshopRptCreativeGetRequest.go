package brandhub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺创意报表数据查询 API请求
taobao.brand.startshop.rpt.creative.get

获取明星店铺广告creative分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStartshopRptCreativeGetRequest struct {
    model.Params
    // 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
    _trafficType   string
    // 当前页数
    _pageIndex   string
    // 每页条数
    _pageSize   string
    // 转化周期,默认15, 3,7,15
    _effect   string
    // 开始日期(最多查询1个月的数据)
    _startDate   string
    // 截至日期(最晚到昨天)
    _endDate   string
}

// 初始化TaobaoBrandStartshopRptCreativeGetRequest对象
func NewTaobaoBrandStartshopRptCreativeGetRequest() *TaobaoBrandStartshopRptCreativeGetRequest{
    return &TaobaoBrandStartshopRptCreativeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBrandStartshopRptCreativeGetRequest) GetApiMethodName() string {
    return "taobao.brand.startshop.rpt.creative.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBrandStartshopRptCreativeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetTrafficType(_trafficType string) error {
    r._trafficType = _trafficType
    r.Set("traffic_type", _trafficType)
    return nil
}

// TrafficType Getter
func (r TaobaoBrandStartshopRptCreativeGetRequest) GetTrafficType() string {
    return r._trafficType
}
// PageIndex Setter
// 当前页数
func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetPageIndex(_pageIndex string) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoBrandStartshopRptCreativeGetRequest) GetPageIndex() string {
    return r._pageIndex
}
// PageSize Setter
// 每页条数
func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetPageSize(_pageSize string) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoBrandStartshopRptCreativeGetRequest) GetPageSize() string {
    return r._pageSize
}
// Effect Setter
// 转化周期,默认15, 3,7,15
func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetEffect(_effect string) error {
    r._effect = _effect
    r.Set("effect", _effect)
    return nil
}

// Effect Getter
func (r TaobaoBrandStartshopRptCreativeGetRequest) GetEffect() string {
    return r._effect
}
// StartDate Setter
// 开始日期(最多查询1个月的数据)
func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoBrandStartshopRptCreativeGetRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 截至日期(最晚到昨天)
func (r *TaobaoBrandStartshopRptCreativeGetRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoBrandStartshopRptCreativeGetRequest) GetEndDate() string {
    return r._endDate
}
