package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeToolsItemsUnbindAPIRequest
交易开放商品解绑 API请求
taobao.opentrade.tools.items.unbind

交易开放商品解绑 */
type TaobaoOpentradeToolsItemsUnbindAPIRequest struct {
	model.Params
	// 绑定交易开放场景的C端小程序ID
	_miniappId int64
	// 商品id
	_itemIds []int64
}

// NewTaobaoOpentradeToolsItemsUnbindRequest 初始化TaobaoOpentradeToolsItemsUnbindAPIRequest对象
func NewTaobaoOpentradeToolsItemsUnbindRequest() *TaobaoOpentradeToolsItemsUnbindAPIRequest {
	return &TaobaoOpentradeToolsItemsUnbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeToolsItemsUnbindAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.tools.items.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeToolsItemsUnbindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MiniappId Setter
// 绑定交易开放场景的C端小程序ID
func (r *TaobaoOpentradeToolsItemsUnbindAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// Get MiniappId Getter
func (r TaobaoOpentradeToolsItemsUnbindAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}

// Set is ItemIds Setter
// 商品id
func (r *TaobaoOpentradeToolsItemsUnbindAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// Get ItemIds Getter
func (r TaobaoOpentradeToolsItemsUnbindAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}
