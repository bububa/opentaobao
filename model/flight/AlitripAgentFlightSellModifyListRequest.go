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
type AlitripAgentFlightSellModifyListRequest struct {
    model.Params
    // 入参
    _param   *ModifyListRequestDto
}

// 初始化AlitripAgentFlightSellModifyListRequest对象
func NewAlitripAgentFlightSellModifyListRequest() *AlitripAgentFlightSellModifyListRequest{
    return &AlitripAgentFlightSellModifyListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyListRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.list"
}

// IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlitripAgentFlightSellModifyListRequest) SetParam(_param *ModifyListRequestDto) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlitripAgentFlightSellModifyListRequest) GetParam() *ModifyListRequestDto {
    return r._param
}
