package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeSpecialItemsUnbindAPIRequest 专属下单场景商品解绑 API请求
// taobao.opentrade.special.items.unbind
//
// 专属下单场景商品解绑
type TaobaoOpentradeSpecialItemsUnbindAPIRequest struct {
	model.Params
	// 本次待解绑的商品ID列表
	_itemIds []int64
	// 绑定专属下单场景的C端小程序ID
	_miniappId int64
}

// NewTaobaoOpentradeSpecialItemsUnbindRequest 初始化TaobaoOpentradeSpecialItemsUnbindAPIRequest对象
func NewTaobaoOpentradeSpecialItemsUnbindRequest() *TaobaoOpentradeSpecialItemsUnbindAPIRequest {
	return &TaobaoOpentradeSpecialItemsUnbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialItemsUnbindAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.special.items.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialItemsUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpentradeSpecialItemsUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 本次待解绑的商品ID列表
func (r *TaobaoOpentradeSpecialItemsUnbindAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoOpentradeSpecialItemsUnbindAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}

// SetMiniappId is MiniappId Setter
// 绑定专属下单场景的C端小程序ID
func (r *TaobaoOpentradeSpecialItemsUnbindAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// GetMiniappId MiniappId Getter
func (r TaobaoOpentradeSpecialItemsUnbindAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}
