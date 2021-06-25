package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道间库存规则设置接口 APIRequest
taobao.qimen.inventoryrule.create

渠道间库存规则设置
*/
type TaobaoQimenInventoryruleCreateRequest struct {
    model.Params

    // 
    request   *RequestDO 

}

func NewTaobaoQimenInventoryruleCreateRequest() *TaobaoQimenInventoryruleCreateRequest{
    return &TaobaoQimenInventoryruleCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenInventoryruleCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.inventoryrule.create"
}

func (r TaobaoQimenInventoryruleCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenInventoryruleCreateRequest) SetRequest(request *RequestDO) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenInventoryruleCreateRequest) GetRequest() *RequestDO {
    return r.request
}

