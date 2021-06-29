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
    applyId   string
    // 国内国际标识
    domesticIntl   int64
    // 拒绝原因
    refuseReason   string
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
func (r *AlitripAgentFlightSellRefundRefuseRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r AlitripAgentFlightSellRefundRefuseRequest) GetApplyId() string {
    return r.applyId
}
// DomesticIntl Setter
// 国内国际标识
func (r *AlitripAgentFlightSellRefundRefuseRequest) SetDomesticIntl(domesticIntl int64) error {
    r.domesticIntl = domesticIntl
    r.Set("domestic_intl", domesticIntl)
    return nil
}

// DomesticIntl Getter
func (r AlitripAgentFlightSellRefundRefuseRequest) GetDomesticIntl() int64 {
    return r.domesticIntl
}
// RefuseReason Setter
// 拒绝原因
func (r *AlitripAgentFlightSellRefundRefuseRequest) SetRefuseReason(refuseReason string) error {
    r.refuseReason = refuseReason
    r.Set("refuse_reason", refuseReason)
    return nil
}

// RefuseReason Getter
func (r AlitripAgentFlightSellRefundRefuseRequest) GetRefuseReason() string {
    return r.refuseReason
}
