package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票拒绝 API请求
alitrip.agent.flight.sell.refund.refuse

销售退票拒绝
*/
type AlitripAgentFlightSellRefundRefuseRequest struct {
    model.Params
    // 申请单号
    _applyId   string
    // 国内国际标识
    _domesticIntl   int64
    // 拒绝原因
    _refuseReason   string
}

// 初始化AlitripAgentFlightSellRefundRefuseRequest对象
func NewAlitripAgentFlightSellRefundRefuseRequest() *AlitripAgentFlightSellRefundRefuseRequest{
    return &AlitripAgentFlightSellRefundRefuseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellRefundRefuseRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.refund.refuse"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellRefundRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 申请单号
func (r *AlitripAgentFlightSellRefundRefuseRequest) SetApplyId(_applyId string) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r AlitripAgentFlightSellRefundRefuseRequest) GetApplyId() string {
    return r._applyId
}
// DomesticIntl Setter
// 国内国际标识
func (r *AlitripAgentFlightSellRefundRefuseRequest) SetDomesticIntl(_domesticIntl int64) error {
    r._domesticIntl = _domesticIntl
    r.Set("domestic_intl", _domesticIntl)
    return nil
}

// DomesticIntl Getter
func (r AlitripAgentFlightSellRefundRefuseRequest) GetDomesticIntl() int64 {
    return r._domesticIntl
}
// RefuseReason Setter
// 拒绝原因
func (r *AlitripAgentFlightSellRefundRefuseRequest) SetRefuseReason(_refuseReason string) error {
    r._refuseReason = _refuseReason
    r.Set("refuse_reason", _refuseReason)
    return nil
}

// RefuseReason Getter
func (r AlitripAgentFlightSellRefundRefuseRequest) GetRefuseReason() string {
    return r._refuseReason
}
