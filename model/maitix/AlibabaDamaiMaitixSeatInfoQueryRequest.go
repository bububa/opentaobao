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
type AlibabaDamaiMaitixSeatInfoQueryAPIRequest struct {
    model.Params
    // 入参
    _seatQueryParam   *SeatQueryParam
}

// 初始化AlibabaDamaiMaitixSeatInfoQueryAPIRequest对象
func NewAlibabaDamaiMaitixSeatInfoQueryRequest() *AlibabaDamaiMaitixSeatInfoQueryAPIRequest{
    return &AlibabaDamaiMaitixSeatInfoQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixSeatInfoQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.seat.info.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixSeatInfoQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SeatQueryParam Setter
// 入参
func (r *AlibabaDamaiMaitixSeatInfoQueryAPIRequest) SetSeatQueryParam(_seatQueryParam *SeatQueryParam) error {
    r._seatQueryParam = _seatQueryParam
    r.Set("seat_query_param", _seatQueryParam)
    return nil
}

// SeatQueryParam Getter
func (r AlibabaDamaiMaitixSeatInfoQueryAPIRequest) GetSeatQueryParam() *SeatQueryParam {
    return r._seatQueryParam
}
