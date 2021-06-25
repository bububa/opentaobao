package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品接口 APIRequest
taobao.qimen.combineitem.synchronize

ERP调用奇门的接口,将商品信息同步给WMS
*/
type TaobaoQimenCombineitemSynchronizeRequest struct {
    model.Params

    // 
    request   *CombineItemSyncRequest 

}

func NewTaobaoQimenCombineitemSynchronizeRequest() *TaobaoQimenCombineitemSynchronizeRequest{
    return &TaobaoQimenCombineitemSynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenCombineitemSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.combineitem.synchronize"
}

func (r TaobaoQimenCombineitemSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenCombineitemSynchronizeRequest) SetRequest(request *CombineItemSyncRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenCombineitemSynchronizeRequest) GetRequest() *CombineItemSyncRequest {
    return r.request
}

