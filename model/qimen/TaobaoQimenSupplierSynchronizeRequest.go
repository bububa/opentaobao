package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商同步接口 APIRequest
taobao.qimen.supplier.synchronize

这个接口用来同步供应商信息
*/
type TaobaoQimenSupplierSynchronizeRequest struct {
    model.Params

    // 
    request   *Request 

}

func NewTaobaoQimenSupplierSynchronizeRequest() *TaobaoQimenSupplierSynchronizeRequest{
    return &TaobaoQimenSupplierSynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenSupplierSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.supplier.synchronize"
}

func (r TaobaoQimenSupplierSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenSupplierSynchronizeRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenSupplierSynchronizeRequest) GetRequest() *Request {
    return r.request
}

