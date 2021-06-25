package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合货品删除接口 APIRequest
taobao.qimen.combineitem.delete

组合货品删除
*/
type TaobaoQimenCombineitemDeleteRequest struct {
    model.Params

    // 
    request   *RequestDO 

}

func NewTaobaoQimenCombineitemDeleteRequest() *TaobaoQimenCombineitemDeleteRequest{
    return &TaobaoQimenCombineitemDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenCombineitemDeleteRequest) GetApiMethodName() string {
    return "taobao.qimen.combineitem.delete"
}

func (r TaobaoQimenCombineitemDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenCombineitemDeleteRequest) SetRequest(request *RequestDO) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenCombineitemDeleteRequest) GetRequest() *RequestDO {
    return r.request
}

