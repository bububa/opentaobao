package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobanamadpcitemeditrenderAPIRequest 编辑商品发布页 API请求
// taobao.banamadpc.item.edit.render
//
// 巴拿马供应商通过此接口获取编辑商品发布页
type TaobaobanamadpcitemeditrenderAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTaobaobanamadpcitemeditrenderRequest 初始化TaobaobanamadpcitemeditrenderAPIRequest对象
func NewTaobaobanamadpcitemeditrenderRequest() *TaobaobanamadpcitemeditrenderAPIRequest {
	return &TaobaobanamadpcitemeditrenderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobanamadpcitemeditrenderAPIRequest) GetApiMethodName() string {
	return "taobao.banamadpc.item.edit.render"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobanamadpcitemeditrenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobanamadpcitemeditrenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaobanamadpcitemeditrenderAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaobanamadpcitemeditrenderAPIRequest) GetItemId() int64 {
	return r._itemId
}
