package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签拒绝 APIRequest
alitrip.agent.flight.sell.modify.refuse

销售改签拒绝
*/
type AlitripAgentFlightSellModifyRefuseRequest struct {
    model.Params

    // 申请单号
    applyId   string 

    // 国际国内标识
    domesticIntl   int64 

    // 拒绝原因
    refuseReason   string 

}

func NewAlitripAgentFlightSellModifyRefuseRequest() *AlitripAgentFlightSellModifyRefuseRequest{
    return &AlitripAgentFlightSellModifyRefuseRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellModifyRefuseRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.refuse"
}

func (r AlitripAgentFlightSellModifyRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellModifyRefuseRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r AlitripAgentFlightSellModifyRefuseRequest) GetApplyId() string {
    return r.applyId
}

func (r *AlitripAgentFlightSellModifyRefuseRequest) SetDomesticIntl(domesticIntl int64) error {
    r.domesticIntl = domesticIntl
    r.Set("domestic_intl", domesticIntl)
    return nil
}

func (r AlitripAgentFlightSellModifyRefuseRequest) GetDomesticIntl() int64 {
    return r.domesticIntl
}

func (r *AlitripAgentFlightSellModifyRefuseRequest) SetRefuseReason(refuseReason string) error {
    r.refuseReason = refuseReason
    r.Set("refuse_reason", refuseReason)
    return nil
}

func (r AlitripAgentFlightSellModifyRefuseRequest) GetRefuseReason() string {
    return r.refuseReason
}

