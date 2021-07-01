package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID获取商品信息 API请求
taobao.wlb.item.get

根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
*/
type TaobaoWlbItemGetAPIRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
}

// 初始化TaobaoWlbItemGetAPIRequest对象
func NewTaobaoWlbItemGetRequest() *TaobaoWlbItemGetAPIRequest{
    return &TaobaoWlbItemGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemGetAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.item.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TaobaoWlbItemGetAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoWlbItemGetAPIRequest) GetItemId() int64 {
    return r._itemId
}
