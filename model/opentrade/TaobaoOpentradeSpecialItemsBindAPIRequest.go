package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeSpecialItemsBindAPIRequest 专属下单场景商品绑定 API请求
// taobao.opentrade.special.items.bind
//
// 专属下单场景商品绑定
type TaobaoOpentradeSpecialItemsBindAPIRequest struct {
	model.Params
	// 本次待绑定的商品ID列表
	_itemIds []int64
	// 绑定专属下单场景的C端小程序ID
	_miniappId int64
}

// NewTaobaoOpentradeSpecialItemsBindRequest 初始化TaobaoOpentradeSpecialItemsBindAPIRequest对象
func NewTaobaoOpentradeSpecialItemsBindRequest() *TaobaoOpentradeSpecialItemsBindAPIRequest {
	return &TaobaoOpentradeSpecialItemsBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.special.items.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemIds is ItemIds Setter
// 本次待绑定的商品ID列表
func (r *TaobaoOpentradeSpecialItemsBindAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}

// SetMiniappId is MiniappId Setter
// 绑定专属下单场景的C端小程序ID
func (r *TaobaoOpentradeSpecialItemsBindAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// GetMiniappId MiniappId Getter
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}
