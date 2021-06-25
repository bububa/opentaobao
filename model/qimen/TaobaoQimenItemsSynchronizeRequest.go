package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品同步接口 (批量) APIRequest
taobao.qimen.items.synchronize

ERP调用奇门的接口,批量同步商品信息给WMS
*/
type TaobaoQimenItemsSynchronizeRequest struct {
    model.Params

    // 
    request   *ItemsSynchronizeRequest 

}

func NewTaobaoQimenItemsSynchronizeRequest() *TaobaoQimenItemsSynchronizeRequest{
    return &TaobaoQimenItemsSynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenItemsSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.items.synchronize"
}

func (r TaobaoQimenItemsSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenItemsSynchronizeRequest) SetRequest(request *ItemsSynchronizeRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenItemsSynchronizeRequest) GetRequest() *ItemsSynchronizeRequest {
    return r.request
}

