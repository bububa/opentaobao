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
type AlitripAgentFlightSellModifyBackfillRequest struct {
    model.Params
    // 入参
    param   *ModifyBackFillRequestDto
}

// 初始化AlitripAgentFlightSellModifyBackfillRequest对象
func NewAlitripAgentFlightSellModifyBackfillRequest() *AlitripAgentFlightSellModifyBackfillRequest{
    return &AlitripAgentFlightSellModifyBackfillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyBackfillRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.backfill"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyBackfillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlitripAgentFlightSellModifyBackfillRequest) SetParam(param *ModifyBackFillRequestDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellModifyBackfillRequest) GetParam() *ModifyBackFillRequestDto {
    return r.param
}
