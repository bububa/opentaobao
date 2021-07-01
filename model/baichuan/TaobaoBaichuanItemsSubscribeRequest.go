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
type TaobaoBaichuanItemsSubscribeAPIRequest struct {
    model.Params
    // 订阅的商品id列表
    _itemIds   []int64
}

// 初始化TaobaoBaichuanItemsSubscribeAPIRequest对象
func NewTaobaoBaichuanItemsSubscribeRequest() *TaobaoBaichuanItemsSubscribeAPIRequest{
    return &TaobaoBaichuanItemsSubscribeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemsSubscribeAPIRequest) GetApiMethodName() string {
    return "taobao.baichuan.items.subscribe"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemsSubscribeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemIds Setter
// 订阅的商品id列表
func (r *TaobaoBaichuanItemsSubscribeAPIRequest) SetItemIds(_itemIds []int64) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoBaichuanItemsSubscribeAPIRequest) GetItemIds() []int64 {
    return r._itemIds
}
