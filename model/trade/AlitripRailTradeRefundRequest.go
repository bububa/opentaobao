package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退票接口 APIRequest
alitrip.rail.trade.refund

退票接口
*/
type AlitripRailTradeRefundRequest struct {
    model.Params

    // 入参
    refundParam   *RefundReq 

}

func NewAlitripRailTradeRefundRequest() *AlitripRailTradeRefundRequest{
    return &AlitripRailTradeRefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripRailTradeRefundRequest) GetApiMethodName() string {
    return "alitrip.rail.trade.refund"
}

func (r AlitripRailTradeRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripRailTradeRefundRequest) SetRefundParam(refundParam *RefundReq) error {
    r.refundParam = refundParam
    r.Set("refund_param", refundParam)
    return nil
}

func (r AlitripRailTradeRefundRequest) GetRefundParam() *RefundReq {
    return r.refundParam
}

