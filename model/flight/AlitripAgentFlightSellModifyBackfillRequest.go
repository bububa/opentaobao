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
    _param   *ModifyBackFillRequestDTO
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
func (r *AlitripAgentFlightSellModifyBackfillRequest) SetParam(_param *ModifyBackFillRequestDTO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellModifyBackfillRequest) GetParam() *ModifyBackFillRequestDTO {
    return r._param
}
