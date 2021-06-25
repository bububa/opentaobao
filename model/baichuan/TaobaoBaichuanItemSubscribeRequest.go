package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个商品订阅 APIRequest
taobao.baichuan.item.subscribe

百川单个商品订阅
*/
type TaobaoBaichuanItemSubscribeRequest struct {
    model.Params

    // 商品id
    itemId   int64 

}

func NewTaobaoBaichuanItemSubscribeRequest() *TaobaoBaichuanItemSubscribeRequest{
    return &TaobaoBaichuanItemSubscribeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanItemSubscribeRequest) GetApiMethodName() string {
    return "taobao.baichuan.item.subscribe"
}

func (r TaobaoBaichuanItemSubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanItemSubscribeRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoBaichuanItemSubscribeRequest) GetItemId() int64 {
    return r.itemId
}

