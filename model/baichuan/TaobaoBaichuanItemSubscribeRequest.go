package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个商品订阅 API请求
taobao.baichuan.item.subscribe

百川单个商品订阅
*/
type TaobaoBaichuanItemSubscribeRequest struct {
    model.Params
    // 商品id
    itemId   int64
}

// 初始化TaobaoBaichuanItemSubscribeRequest对象
func NewTaobaoBaichuanItemSubscribeRequest() *TaobaoBaichuanItemSubscribeRequest{
    return &TaobaoBaichuanItemSubscribeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemSubscribeRequest) GetApiMethodName() string {
    return "taobao.baichuan.item.subscribe"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemSubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoBaichuanItemSubscribeRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoBaichuanItemSubscribeRequest) GetItemId() int64 {
    return r.itemId
}
