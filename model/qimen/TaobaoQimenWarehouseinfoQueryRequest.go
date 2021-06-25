package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
货主仓库资源查询接口 APIRequest
taobao.qimen.warehouseinfo.query

货主仓库资源查询
*/
type TaobaoQimenWarehouseinfoQueryRequest struct {
    model.Params

    // 
    request   *Request 

}

func NewTaobaoQimenWarehouseinfoQueryRequest() *TaobaoQimenWarehouseinfoQueryRequest{
    return &TaobaoQimenWarehouseinfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenWarehouseinfoQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.warehouseinfo.query"
}

func (r TaobaoQimenWarehouseinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenWarehouseinfoQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenWarehouseinfoQueryRequest) GetRequest() *Request {
    return r.request
}

