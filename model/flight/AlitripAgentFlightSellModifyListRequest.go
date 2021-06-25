package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签单列表 APIRequest
alitrip.agent.flight.sell.modify.list

销售改签单列表
*/
type AlitripAgentFlightSellModifyListRequest struct {
    model.Params

    // 入参
    param   *ModifyListRequestDto 

}

func NewAlitripAgentFlightSellModifyListRequest() *AlitripAgentFlightSellModifyListRequest{
    return &AlitripAgentFlightSellModifyListRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripAgentFlightSellModifyListRequest) GetApiMethodName() string {
    return "alitrip.agent.flight.sell.modify.list"
}

func (r AlitripAgentFlightSellModifyListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripAgentFlightSellModifyListRequest) SetParam(param *ModifyListRequestDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlitripAgentFlightSellModifyListRequest) GetParam() *ModifyListRequestDto {
    return r.param
}

