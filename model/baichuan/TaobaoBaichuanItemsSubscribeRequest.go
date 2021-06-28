package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川批量商品订阅 APIRequest
taobao.baichuan.items.subscribe

百川批量添加订阅的商品
*/
type TaobaoBaichuanItemsSubscribeRequest struct {
    model.Params

    // 订阅的商品id列表
    itemIds   []int64 

}

func NewTaobaoBaichuanItemsSubscribeRequest() *TaobaoBaichuanItemsSubscribeRequest{
    return &TaobaoBaichuanItemsSubscribeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanItemsSubscribeRequest) GetApiMethodName() string {
    return "taobao.baichuan.items.subscribe"
}

func (r TaobaoBaichuanItemsSubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanItemsSubscribeRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoBaichuanItemsSubscribeRequest) GetItemIds() []int64 {
    return r.itemIds
}

