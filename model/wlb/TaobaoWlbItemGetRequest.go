package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID获取商品信息 APIRequest
taobao.wlb.item.get

根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
*/
type TaobaoWlbItemGetRequest struct {
    model.Params

    // 商品ID
    itemId   int64 

}

func NewTaobaoWlbItemGetRequest() *TaobaoWlbItemGetRequest{
    return &TaobaoWlbItemGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbItemGetRequest) GetApiMethodName() string {
    return "taobao.wlb.item.get"
}

func (r TaobaoWlbItemGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbItemGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoWlbItemGetRequest) GetItemId() int64 {
    return r.itemId
}

