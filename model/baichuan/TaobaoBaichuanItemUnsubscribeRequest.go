package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个删除订阅关系 API请求
taobao.baichuan.item.unsubscribe

删除单个商品订阅关系
*/
type TaobaoBaichuanItemUnsubscribeRequest struct {
    model.Params
    // 商品id
    _itemId   int64
}

// 初始化TaobaoBaichuanItemUnsubscribeRequest对象
func NewTaobaoBaichuanItemUnsubscribeRequest() *TaobaoBaichuanItemUnsubscribeRequest{
    return &TaobaoBaichuanItemUnsubscribeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemUnsubscribeRequest) GetApiMethodName() string {
    return "taobao.baichuan.item.unsubscribe"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemUnsubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoBaichuanItemUnsubscribeRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoBaichuanItemUnsubscribeRequest) GetItemId() int64 {
    return r._itemId
}
