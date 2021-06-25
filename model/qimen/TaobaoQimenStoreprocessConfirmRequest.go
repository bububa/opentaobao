package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内加工单确认接口 APIRequest
taobao.qimen.storeprocess.confirm

WMS调用奇门的接口,回传仓内加工单创建情况
*/
type TaobaoQimenStoreprocessConfirmRequest struct {
    model.Params

    // 
    request   *StoreProcessConfirmRequest 

}

func NewTaobaoQimenStoreprocessConfirmRequest() *TaobaoQimenStoreprocessConfirmRequest{
    return &TaobaoQimenStoreprocessConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStoreprocessConfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.storeprocess.confirm"
}

func (r TaobaoQimenStoreprocessConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStoreprocessConfirmRequest) SetRequest(request *StoreProcessConfirmRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenStoreprocessConfirmRequest) GetRequest() *StoreProcessConfirmRequest {
    return r.request
}

