package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CRS接送机订单列表搜索 API请求
alitrip.travel.crsorder.search

提供给CRS商家搜索订单列表，仅返回订单号列表
*/
type AlitripTravelCrsorderSearchRequest struct {
    model.Params
    // 订单状态，10-待派单，20-待用车，30-已取消，40-待处理退款申请，60-已关闭，70-已完成
    _crsOrderStatus   int64
    // 用车时间-起始
    _beginCarUseTime   string
    // 页大小，默认20
    _pageSize   int64
    // 用车时间-终止
    _endCarUseTime   string
    // 当前页，默认值1
    _currentPage   int64
    // 支付时间-终止
    _endPayTime   string
    // 支付时间-起始
    _beginPayTime   string
    // 取消时间-起始
    _beginCancelTime   string
    // 取消时间-终止
    _endCancelTime   string
}

// 初始化AlitripTravelCrsorderSearchRequest对象
func NewAlitripTravelCrsorderSearchRequest() *AlitripTravelCrsorderSearchRequest{
    return &AlitripTravelCrsorderSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelCrsorderSearchRequest) GetApiMethodName() string {
    return "alitrip.travel.crsorder.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelCrsorderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CrsOrderStatus Setter
// 订单状态，10-待派单，20-待用车，30-已取消，40-待处理退款申请，60-已关闭，70-已完成
func (r *AlitripTravelCrsorderSearchRequest) SetCrsOrderStatus(_crsOrderStatus int64) error {
    r._crsOrderStatus = _crsOrderStatus
    r.Set("crs_order_status", _crsOrderStatus)
    return nil
}

// CrsOrderStatus Getter
func (r AlitripTravelCrsorderSearchRequest) GetCrsOrderStatus() int64 {
    return r._crsOrderStatus
}
// BeginCarUseTime Setter
// 用车时间-起始
func (r *AlitripTravelCrsorderSearchRequest) SetBeginCarUseTime(_beginCarUseTime string) error {
    r._beginCarUseTime = _beginCarUseTime
    r.Set("begin_car_use_time", _beginCarUseTime)
    return nil
}

// BeginCarUseTime Getter
func (r AlitripTravelCrsorderSearchRequest) GetBeginCarUseTime() string {
    return r._beginCarUseTime
}
// PageSize Setter
// 页大小，默认20
func (r *AlitripTravelCrsorderSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlitripTravelCrsorderSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
// EndCarUseTime Setter
// 用车时间-终止
func (r *AlitripTravelCrsorderSearchRequest) SetEndCarUseTime(_endCarUseTime string) error {
    r._endCarUseTime = _endCarUseTime
    r.Set("end_car_use_time", _endCarUseTime)
    return nil
}

// EndCarUseTime Getter
func (r AlitripTravelCrsorderSearchRequest) GetEndCarUseTime() string {
    return r._endCarUseTime
}
// CurrentPage Setter
// 当前页，默认值1
func (r *AlitripTravelCrsorderSearchRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlitripTravelCrsorderSearchRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// EndPayTime Setter
// 支付时间-终止
func (r *AlitripTravelCrsorderSearchRequest) SetEndPayTime(_endPayTime string) error {
    r._endPayTime = _endPayTime
    r.Set("end_pay_time", _endPayTime)
    return nil
}

// EndPayTime Getter
func (r AlitripTravelCrsorderSearchRequest) GetEndPayTime() string {
    return r._endPayTime
}
// BeginPayTime Setter
// 支付时间-起始
func (r *AlitripTravelCrsorderSearchRequest) SetBeginPayTime(_beginPayTime string) error {
    r._beginPayTime = _beginPayTime
    r.Set("begin_pay_time", _beginPayTime)
    return nil
}

// BeginPayTime Getter
func (r AlitripTravelCrsorderSearchRequest) GetBeginPayTime() string {
    return r._beginPayTime
}
// BeginCancelTime Setter
// 取消时间-起始
func (r *AlitripTravelCrsorderSearchRequest) SetBeginCancelTime(_beginCancelTime string) error {
    r._beginCancelTime = _beginCancelTime
    r.Set("begin_cancel_time", _beginCancelTime)
    return nil
}

// BeginCancelTime Getter
func (r AlitripTravelCrsorderSearchRequest) GetBeginCancelTime() string {
    return r._beginCancelTime
}
// EndCancelTime Setter
// 取消时间-终止
func (r *AlitripTravelCrsorderSearchRequest) SetEndCancelTime(_endCancelTime string) error {
    r._endCancelTime = _endCancelTime
    r.Set("end_cancel_time", _endCancelTime)
    return nil
}

// EndCancelTime Getter
func (r AlitripTravelCrsorderSearchRequest) GetEndCancelTime() string {
    return r._endCancelTime
}
