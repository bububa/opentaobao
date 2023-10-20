package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemvipupdateschemagetAPIRequest vip商家编辑商品的规则获取接口 API请求
// tmall.item.vip.update.schema.get
//
// 获取vip商家编辑商品的规则
type TmallitemvipupdateschemagetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTmallitemvipupdateschemagetRequest 初始化TmallitemvipupdateschemagetAPIRequest对象
func NewTmallitemvipupdateschemagetRequest() *TmallitemvipupdateschemagetAPIRequest {
	return &TmallitemvipupdateschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemvipupdateschemagetAPIRequest) GetApiMethodName() string {
	return "tmall.item.vip.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemvipupdateschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemvipupdateschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallitemvipupdateschemagetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemvipupdateschemagetAPIRequest) GetItemId() int64 {
	return r._itemId
}
