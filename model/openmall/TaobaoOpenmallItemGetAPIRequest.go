package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmallitemgetAPIRequest 获取商品详情物料 API请求
// taobao.openmall.item.get
//
// 获取联盟开放的openmall商品
type TaobaoopenmallitemgetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// NewTaobaoopenmallitemgetRequest 初始化TaobaoopenmallitemgetAPIRequest对象
func NewTaobaoopenmallitemgetRequest() *TaobaoopenmallitemgetAPIRequest {
	return &TaobaoopenmallitemgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmallitemgetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmallitemgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmallitemgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoopenmallitemgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoopenmallitemgetAPIRequest) GetItemId() int64 {
	return r._itemId
}
