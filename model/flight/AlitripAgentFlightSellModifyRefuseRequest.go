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
    _applyId   string
    // 国际国内标识
    _domesticIntl   int64
    // 拒绝原因
    _refuseReason   string
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
func (r *AlitripAgentFlightSellModifyRefuseRequest) SetApplyId(_applyId string) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r AlitripAgentFlightSellModifyRefuseRequest) GetApplyId() string {
    return r._applyId
}
// DomesticIntl Setter
// 国际国内标识
func (r *AlitripAgentFlightSellModifyRefuseRequest) SetDomesticIntl(_domesticIntl int64) error {
    r._domesticIntl = _domesticIntl
    r.Set("domestic_intl", _domesticIntl)
    return nil
}

// DomesticIntl Getter
func (r AlitripAgentFlightSellModifyRefuseRequest) GetDomesticIntl() int64 {
    return r._domesticIntl
}
// RefuseReason Setter
// 拒绝原因
func (r *AlitripAgentFlightSellModifyRefuseRequest) SetRefuseReason(_refuseReason string) error {
    r._refuseReason = _refuseReason
    r.Set("refuse_reason", _refuseReason)
    return nil
}

// RefuseReason Getter
func (r AlitripAgentFlightSellModifyRefuseRequest) GetRefuseReason() string {
    return r._refuseReason
}
