package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易开放商品解绑 API请求
taobao.opentrade.tools.items.unbind

交易开放商品解绑
*/
type TaobaoOpentradeToolsItemsUnbindRequest struct {
    model.Params
    // 绑定交易开放场景的C端小程序ID
    _miniappId   int64
    // 商品id
    _itemIds   []int64
}

// 初始化TaobaoOpentradeToolsItemsUnbindRequest对象
func NewTaobaoOpentradeToolsItemsUnbindRequest() *TaobaoOpentradeToolsItemsUnbindRequest{
    return &TaobaoOpentradeToolsItemsUnbindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeToolsItemsUnbindRequest) GetApiMethodName() string {
    return "taobao.opentrade.tools.items.unbind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeToolsItemsUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniappId Setter
// 绑定交易开放场景的C端小程序ID
func (r *TaobaoOpentradeToolsItemsUnbindRequest) SetMiniappId(_miniappId int64) error {
    r._miniappId = _miniappId
    r.Set("miniapp_id", _miniappId)
    return nil
}

// MiniappId Getter
func (r TaobaoOpentradeToolsItemsUnbindRequest) GetMiniappId() int64 {
    return r._miniappId
}
// ItemIds Setter
// 商品id
func (r *TaobaoOpentradeToolsItemsUnbindRequest) SetItemIds(_itemIds []int64) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoOpentradeToolsItemsUnbindRequest) GetItemIds() []int64 {
    return r._itemIds
}
