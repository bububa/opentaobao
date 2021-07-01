package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除商品订阅 API请求
taobao.baichuan.items.unsubscribe

批量删除商品订阅
*/
type TaobaoBaichuanItemsUnsubscribeAPIRequest struct {
    model.Params
    // 删除的商品id
    _itemIds   []int64
}

// 初始化TaobaoBaichuanItemsUnsubscribeAPIRequest对象
func NewTaobaoBaichuanItemsUnsubscribeRequest() *TaobaoBaichuanItemsUnsubscribeAPIRequest{
    return &TaobaoBaichuanItemsUnsubscribeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemsUnsubscribeAPIRequest) GetApiMethodName() string {
    return "taobao.baichuan.items.unsubscribe"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemsUnsubscribeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemIds Setter
// 删除的商品id
func (r *TaobaoBaichuanItemsUnsubscribeAPIRequest) SetItemIds(_itemIds []int64) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoBaichuanItemsUnsubscribeAPIRequest) GetItemIds() []int64 {
    return r._itemIds
}
