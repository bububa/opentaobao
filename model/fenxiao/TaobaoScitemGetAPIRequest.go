package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoscitemgetAPIRequest 根据id查询商品 API请求
// taobao.scitem.get
//
// 根据id查询商品
type TaobaoscitemgetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTaobaoscitemgetRequest 初始化TaobaoscitemgetAPIRequest对象
func NewTaobaoscitemgetRequest() *TaobaoscitemgetAPIRequest {
	return &TaobaoscitemgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoscitemgetAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoscitemgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoscitemgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoscitemgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoscitemgetAPIRequest) GetItemId() int64 {
	return r._itemId
}
