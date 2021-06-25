package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
前后端商品映射接口 APIRequest
taobao.qimen.itemmapping.create

前后端商品映射
*/
type TaobaoQimenItemmappingCreateRequest struct {
    model.Params

    // 
    request   *RequestDO 

}

func NewTaobaoQimenItemmappingCreateRequest() *TaobaoQimenItemmappingCreateRequest{
    return &TaobaoQimenItemmappingCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenItemmappingCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.itemmapping.create"
}

func (r TaobaoQimenItemmappingCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenItemmappingCreateRequest) SetRequest(request *RequestDO) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenItemmappingCreateRequest) GetRequest() *RequestDO {
    return r.request
}

