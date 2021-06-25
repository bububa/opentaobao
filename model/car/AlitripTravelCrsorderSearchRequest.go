package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CRS接送机订单列表搜索 APIRequest
alitrip.travel.crsorder.search

提供给CRS商家搜索订单列表，仅返回订单号列表
*/
type AlitripTravelCrsorderSearchRequest struct {
    model.Params

    // 订单状态，10-待派单，20-待用车，30-已取消，40-待处理退款申请，60-已关闭，70-已完成
    crsOrderStatus   int64 

    // 用车时间-起始
    beginCarUseTime   string 

    // 页大小，默认20
    pageSize   int64 

    // 用车时间-终止
    endCarUseTime   string 

    // 当前页，默认值1
    currentPage   int64 

    // 支付时间-终止
    endPayTime   string 

    // 支付时间-起始
    beginPayTime   string 

    // 取消时间-起始
    beginCancelTime   string 

    // 取消时间-终止
    endCancelTime   string 

}

func NewAlitripTravelCrsorderSearchRequest() *AlitripTravelCrsorderSearchRequest{
    return &AlitripTravelCrsorderSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelCrsorderSearchRequest) GetApiMethodName() string {
    return "alitrip.travel.crsorder.search"
}

func (r AlitripTravelCrsorderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelCrsorderSearchRequest) SetCrsOrderStatus(crsOrderStatus int64) error {
    r.crsOrderStatus = crsOrderStatus
    r.Set("crs_order_status", crsOrderStatus)
    return nil
}

func (r AlitripTravelCrsorderSearchRequest) GetCrsOrderStatus() int64 {
    return r.crsOrderStatus
}

func (r *AlitripTravelCrsorderSearchRequest) SetBeginCarUseTime(beginCarUseTime string) error {
    r.beginCarUseTime = beginCarUseTime
    r.Set("begin_car_use_time", beginCarUseTime)
    return nil
}

func (r AlitripTravelCrsorderSearchRequest) GetBeginCarUseTime() string {
    return r.beginCarUseTime
}

func (r *AlitripTravelCrsorderSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlitripTravelCrsorderSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlitripTravelCrsorderSearchRequest) SetEndCarUseTime(endCarUseTime string) error {
    r.endCarUseTime = endCarUseTime
    r.Set("end_car_use_time", endCarUseTime)
    return nil
}

func (r AlitripTravelCrsorderSearchRequest) GetEndCarUseTime() string {
    return r.endCarUseTime
}

func (r *AlitripTravelCrsorderSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlitripTravelCrsorderSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlitripTravelCrsorderSearchRequest) SetEndPayTime(endPayTime string) error {
    r.endPayTime = endPayTime
    r.Set("end_pay_time", endPayTime)
    return nil
}

func (r AlitripTravelCrsorderSearchRequest) GetEndPayTime() string {
    return r.endPayTime
}

func (r *AlitripTravelCrsorderSearchRequest) SetBeginPayTime(beginPayTime string) error {
    r.beginPayTime = beginPayTime
    r.Set("begin_pay_time", beginPayTime)
    return nil
}

func (r AlitripTravelCrsorderSearchRequest) GetBeginPayTime() string {
    return r.beginPayTime
}

func (r *AlitripTravelCrsorderSearchRequest) SetBeginCancelTime(beginCancelTime string) error {
    r.beginCancelTime = beginCancelTime
    r.Set("begin_cancel_time", beginCancelTime)
    return nil
}

func (r AlitripTravelCrsorderSearchRequest) GetBeginCancelTime() string {
    return r.beginCancelTime
}

func (r *AlitripTravelCrsorderSearchRequest) SetEndCancelTime(endCancelTime string) error {
    r.endCancelTime = endCancelTime
    r.Set("end_cancel_time", endCancelTime)
    return nil
}

func (r AlitripTravelCrsorderSearchRequest) GetEndCancelTime() string {
    return r.endCancelTime
}

