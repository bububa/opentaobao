package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票单详情 APIRequest
alitrip.agent.flight.sell.refund.detail

销售退票单详情
*/
type AlitripAgentFlightSellRefundDetailRequest struct {
    model.Params

    // 申请单号
    applyId   string 

    // 国际国内标识
    domesticIntl   int64 

}

func NewAlitripAgentFlightSellRefundDetailRequest() *AlitripAgentFlightSellRefundDetailRequest{
    return &AlitripAgentFlightSellRefundDetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellRefundDetailRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.refund.detail"
}

func (r AlitripAgentFlightSellRefundDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellRefundDetailRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r AlitripAgentFlightSellRefundDetailRequest) GetApplyId() string {
    return r.applyId
}

func (r *AlitripAgentFlightSellRefundDetailRequest) SetDomesticIntl(domesticIntl int64) error {
    r.domesticIntl = domesticIntl
    r.Set("domestic_intl", domesticIntl)
    return nil
}

func (r AlitripAgentFlightSellRefundDetailRequest) GetDomesticIntl() int64 {
    return r.domesticIntl
}

