package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemUpdateSimpleschemaGetAPIRequest
官网同购编辑商品的get接口 API请求
tmall.item.update.simpleschema.get

官网同购编辑商品的get接口 */
type TmallItemUpdateSimpleschemaGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTmallItemUpdateSimpleschemaGetRequest 初始化TmallItemUpdateSimpleschemaGetAPIRequest对象
func NewTmallItemUpdateSimpleschemaGetRequest() *TmallItemUpdateSimpleschemaGetAPIRequest {
	return &TmallItemUpdateSimpleschemaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemUpdateSimpleschemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.update.simpleschema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemUpdateSimpleschemaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 商品id
func (r *TmallItemUpdateSimpleschemaGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TmallItemUpdateSimpleschemaGetAPIRequest) GetItemId() int64 {
	return r._itemId
}
