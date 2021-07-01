package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签确认 API请求
alitrip.agent.flight.sell.modify.approve

销售改签确认
*/
type AlitripAgentFlightSellModifyApproveAPIRequest struct {
    model.Params
    // 入参对象
    _param   *ModifyApproveRequestDto
}

// 初始化AlitripAgentFlightSellModifyApproveAPIRequest对象
func NewAlitripAgentFlightSellModifyApproveRequest() *AlitripAgentFlightSellModifyApproveAPIRequest{
    return &AlitripAgentFlightSellModifyApproveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyApproveAPIRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.approve"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyApproveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参对象
func (r *AlitripAgentFlightSellModifyApproveAPIRequest) SetParam(_param *ModifyApproveRequestDto) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellModifyApproveAPIRequest) GetParam() *ModifyApproveRequestDto {
    return r._param
}
