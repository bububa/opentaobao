package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品同步接口 APIRequest
taobao.qimen.singleitem.synchronize

ERP调用奇门的接口,同步商品信息给WMS
*/
type TaobaoQimenSingleitemSynchronizeRequest struct {
    model.Params

    // 
    request   *ItemSynRequest 

}

func NewTaobaoQimenSingleitemSynchronizeRequest() *TaobaoQimenSingleitemSynchronizeRequest{
    return &TaobaoQimenSingleitemSynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenSingleitemSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.singleitem.synchronize"
}

func (r TaobaoQimenSingleitemSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenSingleitemSynchronizeRequest) SetRequest(request *ItemSynRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenSingleitemSynchronizeRequest) GetRequest() *ItemSynRequest {
    return r.request
}

