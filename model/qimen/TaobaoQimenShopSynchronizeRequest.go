package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
店铺同步接口 APIRequest
taobao.qimen.shop.synchronize

店铺同步接口描述
*/
type TaobaoQimenShopSynchronizeRequest struct {
    model.Params

    // 请求
    request   *Request 

}

func NewTaobaoQimenShopSynchronizeRequest() *TaobaoQimenShopSynchronizeRequest{
    return &TaobaoQimenShopSynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenShopSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.shop.synchronize"
}

func (r TaobaoQimenShopSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenShopSynchronizeRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenShopSynchronizeRequest) GetRequest() *Request {
    return r.request
}

