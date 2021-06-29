package brandhub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺推广单元报表数据查询 API请求
taobao.brand.startshop.rpt.adgroup.get

获取明星店铺广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStartshopRptAdgroupGetRequest struct {
    model.Params
    // 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
    _trafficType   string
    // 转化周期默认15天,3,7,15
    _effect   int64
    // 当前页数
    _pageIndex   string
    // 每页条数
    _pageSize   string
    // 开始时间(最多可查询最近90天)
    _startDate   string
    // 截至时间(最晚到昨天)
    _endDate   string
}

// 初始化TaobaoBrandStartshopRptAdgroupGetRequest对象
func NewTaobaoBrandStartshopRptAdgroupGetRequest() *TaobaoBrandStartshopRptAdgroupGetRequest{
    return &TaobaoBrandStartshopRptAdgroupGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetApiMethodName() string {
    return "taobao.brand.startshop.rpt.adgroup.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetTrafficType(_trafficType string) error {
    r._trafficType = _trafficType
    r.Set("traffic_type", _trafficType)
    return nil
}

// TrafficType Getter
func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetTrafficType() string {
    return r._trafficType
}
// Effect Setter
// 转化周期默认15天,3,7,15
func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetEffect(_effect int64) error {
    r._effect = _effect
    r.Set("effect", _effect)
    return nil
}

// Effect Getter
func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetEffect() int64 {
    return r._effect
}
// PageIndex Setter
// 当前页数
func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetPageIndex(_pageIndex string) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetPageIndex() string {
    return r._pageIndex
}
// PageSize Setter
// 每页条数
func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetPageSize(_pageSize string) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetPageSize() string {
    return r._pageSize
}
// StartDate Setter
// 开始时间(最多可查询最近90天)
func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 截至时间(最晚到昨天)
func (r *TaobaoBrandStartshopRptAdgroupGetRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoBrandStartshopRptAdgroupGetRequest) GetEndDate() string {
    return r._endDate
}