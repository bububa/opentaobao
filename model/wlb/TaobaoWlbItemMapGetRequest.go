package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据物流宝商品ID查询商品映射关系 APIRequest
taobao.wlb.item.map.get

根据物流宝商品ID查询商品映射关系
*/
type TaobaoWlbItemMapGetRequest struct {
    model.Params

    // 要查询映射关系的物流宝商品id
    itemId   int64 

}

func NewTaobaoWlbItemMapGetRequest() *TaobaoWlbItemMapGetRequest{
    return &TaobaoWlbItemMapGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbItemMapGetRequest) GetApiMethodName() string {
    return "taobao.wlb.item.map.get"
}

func (r TaobaoWlbItemMapGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbItemMapGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoWlbItemMapGetRequest) GetItemId() int64 {
    return r.itemId
}

