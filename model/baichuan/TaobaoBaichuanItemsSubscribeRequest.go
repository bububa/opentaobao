package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川批量商品订阅 API请求
taobao.baichuan.items.subscribe

百川批量添加订阅的商品
*/
type TaobaoBaichuanItemsSubscribeRequest struct {
    model.Params
    // 订阅的商品id列表
    itemIds   []int64
}

// 初始化TaobaoBaichuanItemsSubscribeRequest对象
func NewTaobaoBaichuanItemsSubscribeRequest() *TaobaoBaichuanItemsSubscribeRequest{
    return &TaobaoBaichuanItemsSubscribeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemsSubscribeRequest) GetApiMethodName() string {
    return "taobao.baichuan.items.subscribe"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemsSubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemIds Setter
// 订阅的商品id列表
func (r *TaobaoBaichuanItemsSubscribeRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoBaichuanItemsSubscribeRequest) GetItemIds() []int64 {
    return r.itemIds
}
