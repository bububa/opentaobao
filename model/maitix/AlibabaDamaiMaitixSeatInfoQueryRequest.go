package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商查询座位信息 APIRequest
alibaba.damai.maitix.seat.info.query

分销查询座位文案信息
*/
type AlibabaDamaiMaitixSeatInfoQueryRequest struct {
    model.Params

    // 入参
    seatQueryParam   *SeatQueryParam 

}

func NewAlibabaDamaiMaitixSeatInfoQueryRequest() *AlibabaDamaiMaitixSeatInfoQueryRequest{
    return &AlibabaDamaiMaitixSeatInfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixSeatInfoQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.seat.info.query"
}

func (r AlibabaDamaiMaitixSeatInfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixSeatInfoQueryRequest) SetSeatQueryParam(seatQueryParam *SeatQueryParam) error {
    r.seatQueryParam = seatQueryParam
    r.Set("seat_query_param", seatQueryParam)
    return nil
}

func (r AlibabaDamaiMaitixSeatInfoQueryRequest) GetSeatQueryParam() *SeatQueryParam {
    return r.seatQueryParam
}

