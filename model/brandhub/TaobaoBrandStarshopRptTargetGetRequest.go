package brandhub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺定向维度报表 API请求
taobao.brand.starshop.rpt.target.get

获取明星店铺定向维度分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStarshopRptTargetGetRequest struct {
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

// 初始化TaobaoBrandStarshopRptTargetGetRequest对象
func NewTaobaoBrandStarshopRptTargetGetRequest() *TaobaoBrandStarshopRptTargetGetRequest{
    return &TaobaoBrandStarshopRptTargetGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBrandStarshopRptTargetGetRequest) GetApiMethodName() string {
    return "taobao.brand.starshop.rpt.target.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBrandStarshopRptTargetGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoBrandStarshopRptTargetGetRequest) SetTrafficType(_trafficType string) error {
    r._trafficType = _trafficType
    r.Set("traffic_type", _trafficType)
    return nil
}

// TrafficType Getter
func (r TaobaoBrandStarshopRptTargetGetRequest) GetTrafficType() string {
    return r._trafficType
}
// PageIndex Setter
// 当前页数
func (r *TaobaoBrandStarshopRptTargetGetRequest) SetPageIndex(_pageIndex string) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoBrandStarshopRptTargetGetRequest) GetPageIndex() string {
    return r._pageIndex
}
// PageSize Setter
// 每页条数
func (r *TaobaoBrandStarshopRptTargetGetRequest) SetPageSize(_pageSize string) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoBrandStarshopRptTargetGetRequest) GetPageSize() string {
    return r._pageSize
}
// Effect Setter
// 转化周期,默认15, 3,7,15
func (r *TaobaoBrandStarshopRptTargetGetRequest) SetEffect(_effect string) error {
    r._effect = _effect
    r.Set("effect", _effect)
    return nil
}

// Effect Getter
func (r TaobaoBrandStarshopRptTargetGetRequest) GetEffect() string {
    return r._effect
}
// StartDate Setter
// 开始日期(最多查询1个月的数据)
func (r *TaobaoBrandStarshopRptTargetGetRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoBrandStarshopRptTargetGetRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 截至日期(最晚到昨天)
func (r *TaobaoBrandStarshopRptTargetGetRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoBrandStarshopRptTargetGetRequest) GetEndDate() string {
    return r._endDate
}
