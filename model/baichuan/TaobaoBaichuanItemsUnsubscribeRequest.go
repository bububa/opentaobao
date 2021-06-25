package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除商品订阅 APIRequest
taobao.baichuan.items.unsubscribe

批量删除商品订阅
*/
type TaobaoBaichuanItemsUnsubscribeRequest struct {
    model.Params

    // 删除的商品id
    itemIds   []Number 

}

func NewTaobaoBaichuanItemsUnsubscribeRequest() *TaobaoBaichuanItemsUnsubscribeRequest{
    return &TaobaoBaichuanItemsUnsubscribeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanItemsUnsubscribeRequest) GetApiMethodName() string {
    return "taobao.baichuan.items.unsubscribe"
}

func (r TaobaoBaichuanItemsUnsubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanItemsUnsubscribeRequest) SetItemIds(itemIds []Number) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoBaichuanItemsUnsubscribeRequest) GetItemIds() []Number {
    return r.itemIds
}

