package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品id查询商品组合关系 APIRequest
taobao.wlb.item.combination.get

根据商品id查询商品组合关系
*/
type TaobaoWlbItemCombinationGetRequest struct {
    model.Params

    // 要查询的组合商品id
    itemId   int64 

}

func NewTaobaoWlbItemCombinationGetRequest() *TaobaoWlbItemCombinationGetRequest{
    return &TaobaoWlbItemCombinationGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbItemCombinationGetRequest) GetApiMethodName() string {
    return "taobao.wlb.item.combination.get"
}

func (r TaobaoWlbItemCombinationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbItemCombinationGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoWlbItemCombinationGetRequest) GetItemId() int64 {
    return r.itemId
}

