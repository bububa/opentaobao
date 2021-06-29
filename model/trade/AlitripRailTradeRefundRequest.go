package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退票接口 API请求
alitrip.rail.trade.refund

退票接口
*/
type AlitripRailTradeRefundRequest struct {
    model.Params
    // 入参
    _refundParam   *RefundReq
}

// 初始化AlitripRailTradeRefundRequest对象
func NewAlitripRailTradeRefundRequest() *AlitripRailTradeRefundRequest{
    return &AlitripRailTradeRefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripRailTradeRefundRequest) GetApiMethodName() string {
    return "alitrip.rail.trade.refund"
}

// IRequest interface 方法, 获取API参数
func (r AlitripRailTradeRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundParam Setter
// 入参
func (r *AlitripRailTradeRefundRequest) SetRefundParam(_refundParam *RefundReq) error {
    r._refundParam = _refundParam
    r.Set("refund_param", _refundParam)
    return nil
}

// RefundParam Getter
func (r AlitripRailTradeRefundRequest) GetRefundParam() *RefundReq {
    return r._refundParam
}
