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
type TaobaoBaichuanItemsUnsubscribeRequest struct {
    model.Params
    // 删除的商品id
    itemIds   []int64
}

// 初始化TaobaoBaichuanItemsUnsubscribeRequest对象
func NewTaobaoBaichuanItemsUnsubscribeRequest() *TaobaoBaichuanItemsUnsubscribeRequest{
    return &TaobaoBaichuanItemsUnsubscribeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemsUnsubscribeRequest) GetApiMethodName() string {
    return "taobao.baichuan.items.unsubscribe"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemsUnsubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemIds Setter
// 删除的商品id
func (r *TaobaoBaichuanItemsUnsubscribeRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoBaichuanItemsUnsubscribeRequest) GetItemIds() []int64 {
    return r.itemIds
}
