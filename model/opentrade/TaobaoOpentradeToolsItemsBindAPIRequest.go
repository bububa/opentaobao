package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeToolsItemsBindAPIRequest 交易开放商品绑定 API请求
// taobao.opentrade.tools.items.bind
//
// 交易开放商品绑定
type TaobaoOpentradeToolsItemsBindAPIRequest struct {
	model.Params
	// 待绑定商品id
	_itemIds []int64
	// 绑定交易开放场景的C端小程序ID
	_miniappId int64
}

// NewTaobaoOpentradeToolsItemsBindRequest 初始化TaobaoOpentradeToolsItemsBindAPIRequest对象
func NewTaobaoOpentradeToolsItemsBindRequest() *TaobaoOpentradeToolsItemsBindAPIRequest {
	return &TaobaoOpentradeToolsItemsBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeToolsItemsBindAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.tools.items.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeToolsItemsBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemIds is ItemIds Setter
// 待绑定商品id
func (r *TaobaoOpentradeToolsItemsBindAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoOpentradeToolsItemsBindAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}

// SetMiniappId is MiniappId Setter
// 绑定交易开放场景的C端小程序ID
func (r *TaobaoOpentradeToolsItemsBindAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// GetMiniappId MiniappId Getter
func (r TaobaoOpentradeToolsItemsBindAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}
