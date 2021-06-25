package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品查询接口 APIRequest
taobao.qimen.singleitem.query

商品查询接口
*/
type TaobaoQimenSingleitemQueryRequest struct {
    model.Params

    // 
    request   *RequestDO 

}

func NewTaobaoQimenSingleitemQueryRequest() *TaobaoQimenSingleitemQueryRequest{
    return &TaobaoQimenSingleitemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenSingleitemQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.singleitem.query"
}

func (r TaobaoQimenSingleitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenSingleitemQueryRequest) SetRequest(request *RequestDO) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenSingleitemQueryRequest) GetRequest() *RequestDO {
    return r.request
}

