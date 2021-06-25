package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道库存查询接口 APIRequest
taobao.qimen.channelinventory.query

渠道库存查询
*/
type TaobaoQimenChannelinventoryQueryRequest struct {
    model.Params

    // 
    request   *RequestDO 

}

func NewTaobaoQimenChannelinventoryQueryRequest() *TaobaoQimenChannelinventoryQueryRequest{
    return &TaobaoQimenChannelinventoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenChannelinventoryQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.channelinventory.query"
}

func (r TaobaoQimenChannelinventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenChannelinventoryQueryRequest) SetRequest(request *RequestDO) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenChannelinventoryQueryRequest) GetRequest() *RequestDO {
    return r.request
}

