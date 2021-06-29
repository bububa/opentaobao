package idleitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
免费送商品发送 APIRequest
alibaba.idle.item.idlecoin.add

免费送商品发布
*/
type AlibabaIdleItemIdlecoinAddRequest struct {
    model.Params

    // 免费送商品数据
    idleCoinItemParam   *IdleCoinItemApiDto 

}

func NewAlibabaIdleItemIdlecoinAddRequest() *AlibabaIdleItemIdlecoinAddRequest{
    return &AlibabaIdleItemIdlecoinAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleItemIdlecoinAddRequest) GetApiMethodName() string {
    return "alibaba.idle.item.idlecoin.add"
}

func (r AlibabaIdleItemIdlecoinAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleItemIdlecoinAddRequest) SetIdleCoinItemParam(idleCoinItemParam *IdleCoinItemApiDto) error {
    r.idleCoinItemParam = idleCoinItemParam
    r.Set("idle_coin_item_param", idleCoinItemParam)
    return nil
}

func (r AlibabaIdleItemIdlecoinAddRequest) GetIdleCoinItemParam() *IdleCoinItemApiDto {
    return r.idleCoinItemParam
}

