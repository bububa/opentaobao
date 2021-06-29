package tuike

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推客网盟合作抽佣订单查询接口 API请求
alibaba.tuike.web.union.order.query

推客网盟合作抽佣订单查询接口
*/
type AlibabaTuikeWebUnionOrderQueryRequest struct {
    model.Params
    // 0 表示time为下单时间;1表示time为更新时间
    timeType   int64
    // 13位时间戳
    startTime   int64
    // 13位时间戳
    endTime   int64
    // 页码偏移
    pageOffset   int64
    // 返回条数
    pageSize   int64
}

// 初始化AlibabaTuikeWebUnionOrderQueryRequest对象
func NewAlibabaTuikeWebUnionOrderQueryRequest() *AlibabaTuikeWebUnionOrderQueryRequest{
    return &AlibabaTuikeWebUnionOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTuikeWebUnionOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.tuike.web.union.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTuikeWebUnionOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TimeType Setter
// 0 表示time为下单时间;1表示time为更新时间
func (r *AlibabaTuikeWebUnionOrderQueryRequest) SetTimeType(timeType int64) error {
    r.timeType = timeType
    r.Set("time_type", timeType)
    return nil
}

// TimeType Getter
func (r AlibabaTuikeWebUnionOrderQueryRequest) GetTimeType() int64 {
    return r.timeType
}
// StartTime Setter
// 13位时间戳
func (r *AlibabaTuikeWebUnionOrderQueryRequest) SetStartTime(startTime int64) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AlibabaTuikeWebUnionOrderQueryRequest) GetStartTime() int64 {
    return r.startTime
}
// EndTime Setter
// 13位时间戳
func (r *AlibabaTuikeWebUnionOrderQueryRequest) SetEndTime(endTime int64) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlibabaTuikeWebUnionOrderQueryRequest) GetEndTime() int64 {
    return r.endTime
}
// PageOffset Setter
// 页码偏移
func (r *AlibabaTuikeWebUnionOrderQueryRequest) SetPageOffset(pageOffset int64) error {
    r.pageOffset = pageOffset
    r.Set("page_offset", pageOffset)
    return nil
}

// PageOffset Getter
func (r AlibabaTuikeWebUnionOrderQueryRequest) GetPageOffset() int64 {
    return r.pageOffset
}
// PageSize Setter
// 返回条数
func (r *AlibabaTuikeWebUnionOrderQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaTuikeWebUnionOrderQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
