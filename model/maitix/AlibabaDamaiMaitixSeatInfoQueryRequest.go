package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商查询座位信息 API请求
alibaba.damai.maitix.seat.info.query

分销查询座位文案信息
*/
type AlibabaDamaiMaitixSeatInfoQueryRequest struct {
    model.Params
    // 入参
    seatQueryParam   *SeatQueryParam
}

// 初始化AlibabaDamaiMaitixSeatInfoQueryRequest对象
func NewAlibabaDamaiMaitixSeatInfoQueryRequest() *AlibabaDamaiMaitixSeatInfoQueryRequest{
    return &AlibabaDamaiMaitixSeatInfoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixSeatInfoQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.seat.info.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixSeatInfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SeatQueryParam Setter
// 入参
func (r *AlibabaDamaiMaitixSeatInfoQueryRequest) SetSeatQueryParam(seatQueryParam *SeatQueryParam) error {
    r.seatQueryParam = seatQueryParam
    r.Set("seat_query_param", seatQueryParam)
    return nil
}

// SeatQueryParam Getter
func (r AlibabaDamaiMaitixSeatInfoQueryRequest) GetSeatQueryParam() *SeatQueryParam {
    return r.seatQueryParam
}
