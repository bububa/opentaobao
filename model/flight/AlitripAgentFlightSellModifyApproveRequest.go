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
type AlitripAgentFlightSellModifyApproveRequest struct {
    model.Params
    // 入参对象
    param   *ModifyApproveRequestDto
}

// 初始化AlitripAgentFlightSellModifyApproveRequest对象
func NewAlitripAgentFlightSellModifyApproveRequest() *AlitripAgentFlightSellModifyApproveRequest{
    return &AlitripAgentFlightSellModifyApproveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyApproveRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.approve"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyApproveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参对象
func (r *AlitripAgentFlightSellModifyApproveRequest) SetParam(param *ModifyApproveRequestDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellModifyApproveRequest) GetParam() *ModifyApproveRequestDto {
    return r.param
}
