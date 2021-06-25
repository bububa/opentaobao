package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内加工单创建接口 APIRequest
taobao.qimen.storeprocess.create

ERP调用奇门的接口,创建仓内加工单
*/
type TaobaoQimenStoreprocessCreateRequest struct {
    model.Params

    // 
    request   *StoreProcessCreateRequest 

}

func NewTaobaoQimenStoreprocessCreateRequest() *TaobaoQimenStoreprocessCreateRequest{
    return &TaobaoQimenStoreprocessCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStoreprocessCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.storeprocess.create"
}

func (r TaobaoQimenStoreprocessCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStoreprocessCreateRequest) SetRequest(request *StoreProcessCreateRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenStoreprocessCreateRequest) GetRequest() *StoreProcessCreateRequest {
    return r.request
}

