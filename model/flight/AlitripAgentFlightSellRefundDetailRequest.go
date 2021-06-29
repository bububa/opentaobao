package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票单详情 API请求
alitrip.agent.flight.sell.refund.detail

销售退票单详情
*/
type AlitripAgentFlightSellRefundDetailRequest struct {
    model.Params
    // 申请单号
    _applyId   string
    // 国际国内标识
    _domesticIntl   int64
}

// 初始化AlitripAgentFlightSellRefundDetailRequest对象
func NewAlitripAgentFlightSellRefundDetailRequest() *AlitripAgentFlightSellRefundDetailRequest{
    return &AlitripAgentFlightSellRefundDetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellRefundDetailRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.refund.detail"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellRefundDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 申请单号
func (r *AlitripAgentFlightSellRefundDetailRequest) SetApplyId(_applyId string) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r AlitripAgentFlightSellRefundDetailRequest) GetApplyId() string {
    return r._applyId
}
// DomesticIntl Setter
// 国际国内标识
func (r *AlitripAgentFlightSellRefundDetailRequest) SetDomesticIntl(_domesticIntl int64) error {
    r._domesticIntl = _domesticIntl
    r.Set("domestic_intl", _domesticIntl)
    return nil
}

// DomesticIntl Getter
func (r AlitripAgentFlightSellRefundDetailRequest) GetDomesticIntl() int64 {
    return r._domesticIntl
}
