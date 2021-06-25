package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票拒绝 APIRequest
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

func NewAlitripAgentFlightSellRefundRefuseRequest() *AlitripAgentFlightSellRefundRefuseRequest{
    return &AlitripAgentFlightSellRefundRefuseRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellRefundRefuseRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.refund.refuse"
}

func (r AlitripAgentFlightSellRefundRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellRefundRefuseRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r AlitripAgentFlightSellRefundRefuseRequest) GetApplyId() string {
    return r.applyId
}

func (r *AlitripAgentFlightSellRefundRefuseRequest) SetDomesticIntl(domesticIntl int64) error {
    r.domesticIntl = domesticIntl
    r.Set("domestic_intl", domesticIntl)
    return nil
}

func (r AlitripAgentFlightSellRefundRefuseRequest) GetDomesticIntl() int64 {
    return r.domesticIntl
}

func (r *AlitripAgentFlightSellRefundRefuseRequest) SetRefuseReason(refuseReason string) error {
    r.refuseReason = refuseReason
    r.Set("refuse_reason", refuseReason)
    return nil
}

func (r AlitripAgentFlightSellRefundRefuseRequest) GetRefuseReason() string {
    return r.refuseReason
}

