package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个删除订阅关系 APIRequest
taobao.baichuan.item.unsubscribe

删除单个商品订阅关系
*/
type TaobaoBaichuanItemUnsubscribeRequest struct {
    model.Params

    // 商品id
    itemId   int64 

}

func NewTaobaoBaichuanItemUnsubscribeRequest() *TaobaoBaichuanItemUnsubscribeRequest{
    return &TaobaoBaichuanItemUnsubscribeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanItemUnsubscribeRequest) GetApiMethodName() string {
    return "taobao.baichuan.item.unsubscribe"
}

func (r TaobaoBaichuanItemUnsubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanItemUnsubscribeRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoBaichuanItemUnsubscribeRequest) GetItemId() int64 {
    return r.itemId
}

