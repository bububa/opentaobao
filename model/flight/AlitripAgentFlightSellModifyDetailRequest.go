package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签详情 APIRequest
alitrip.agent.flight.sell.modify.detail

销售改签详情
*/
type AlitripAgentFlightSellModifyDetailRequest struct {
    model.Params

    // 申请单号
    applyId   string 

    // 国内国际标识
    domesticIntl   int64 

}

func NewAlitripAgentFlightSellModifyDetailRequest() *AlitripAgentFlightSellModifyDetailRequest{
    return &AlitripAgentFlightSellModifyDetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellModifyDetailRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.detail"
}

func (r AlitripAgentFlightSellModifyDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellModifyDetailRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r AlitripAgentFlightSellModifyDetailRequest) GetApplyId() string {
    return r.applyId
}

func (r *AlitripAgentFlightSellModifyDetailRequest) SetDomesticIntl(domesticIntl int64) error {
    r.domesticIntl = domesticIntl
    r.Set("domestic_intl", domesticIntl)
    return nil
}

func (r AlitripAgentFlightSellModifyDetailRequest) GetDomesticIntl() int64 {
    return r.domesticIntl
}

