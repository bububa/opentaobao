package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签单列表 API请求
alitrip.agent.flight.sell.modify.list

销售改签单列表
*/
type AlitripAgentFlightSellModifyListAPIRequest struct {
    model.Params
    // 入参
    _param   *ModifyListRequestDTO
}

// 初始化AlitripAgentFlightSellModifyListAPIRequest对象
func NewAlitripAgentFlightSellModifyListRequest() *AlitripAgentFlightSellModifyListAPIRequest{
    return &AlitripAgentFlightSellModifyListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyListAPIRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.list"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlitripAgentFlightSellModifyListAPIRequest) SetParam(_param *ModifyListRequestDTO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellModifyListAPIRequest) GetParam() *ModifyListRequestDTO {
    return r._param
}
