package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签拒绝 API请求
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

// 初始化AlitripAgentFlightSellModifyRefuseRequest对象
func NewAlitripAgentFlightSellModifyRefuseRequest() *AlitripAgentFlightSellModifyRefuseRequest{
    return &AlitripAgentFlightSellModifyRefuseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyRefuseRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.refuse"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 申请单号
func (r *AlitripAgentFlightSellModifyRefuseRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r AlitripAgentFlightSellModifyRefuseRequest) GetApplyId() string {
    return r.applyId
}
// DomesticIntl Setter
// 国际国内标识
func (r *AlitripAgentFlightSellModifyRefuseRequest) SetDomesticIntl(domesticIntl int64) error {
    r.domesticIntl = domesticIntl
    r.Set("domestic_intl", domesticIntl)
    return nil
}

// DomesticIntl Getter
func (r AlitripAgentFlightSellModifyRefuseRequest) GetDomesticIntl() int64 {
    return r.domesticIntl
}
// RefuseReason Setter
// 拒绝原因
func (r *AlitripAgentFlightSellModifyRefuseRequest) SetRefuseReason(refuseReason string) error {
    r.refuseReason = refuseReason
    r.Set("refuse_reason", refuseReason)
    return nil
}

// RefuseReason Getter
func (r AlitripAgentFlightSellModifyRefuseRequest) GetRefuseReason() string {
    return r.refuseReason
}
