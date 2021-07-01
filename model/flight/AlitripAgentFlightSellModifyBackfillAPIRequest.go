package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签回填 API请求
alitrip.agent.flight.sell.modify.backfill

销售改签回填
*/
type AlitripAgentFlightSellModifyBackfillAPIRequest struct {
    model.Params
    // 入参
    _param   *ModifyBackFillRequestDto
}

// 初始化AlitripAgentFlightSellModifyBackfillAPIRequest对象
func NewAlitripAgentFlightSellModifyBackfillRequest() *AlitripAgentFlightSellModifyBackfillAPIRequest{
    return &AlitripAgentFlightSellModifyBackfillAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyBackfillAPIRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.backfill"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyBackfillAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlitripAgentFlightSellModifyBackfillAPIRequest) SetParam(_param *ModifyBackFillRequestDto) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellModifyBackfillAPIRequest) GetParam() *ModifyBackFillRequestDto {
    return r._param
}
