package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单场景商品解绑 API请求
taobao.opentrade.special.items.unbind

专属下单场景商品解绑
*/
type TaobaoOpentradeSpecialItemsUnbindAPIRequest struct {
    model.Params
    // 绑定专属下单场景的C端小程序ID
    _miniappId   int64
    // 本次待解绑的商品ID列表
    _itemIds   []int64
}

// 初始化TaobaoOpentradeSpecialItemsUnbindAPIRequest对象
func NewTaobaoOpentradeSpecialItemsUnbindRequest() *TaobaoOpentradeSpecialItemsUnbindAPIRequest{
    return &TaobaoOpentradeSpecialItemsUnbindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialItemsUnbindAPIRequest) GetApiMethodName() string {
    return "taobao.opentrade.special.items.unbind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialItemsUnbindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniappId Setter
// 绑定专属下单场景的C端小程序ID
func (r *TaobaoOpentradeSpecialItemsUnbindAPIRequest) SetMiniappId(_miniappId int64) error {
    r._miniappId = _miniappId
    r.Set("miniapp_id", _miniappId)
    return nil
}

// MiniappId Getter
func (r TaobaoOpentradeSpecialItemsUnbindAPIRequest) GetMiniappId() int64 {
    return r._miniappId
}
// ItemIds Setter
// 本次待解绑的商品ID列表
func (r *TaobaoOpentradeSpecialItemsUnbindAPIRequest) SetItemIds(_itemIds []int64) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoOpentradeSpecialItemsUnbindAPIRequest) GetItemIds() []int64 {
    return r._itemIds
}
