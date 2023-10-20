package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenitemstorequeryAPIRequest 商品关联门店查询接口 API请求
// taobao.qimen.itemstore.query
//
// 商家在ERP等系统中调用该接口，查询线上商品所关联的门店列表
type TaobaoqimenitemstorequeryAPIRequest struct {
	model.Params
	// 当前查询的页面编码
	_page int64
	// 线上商品ID
	_itemId int64
}

// NewTaobaoqimenitemstorequeryRequest 初始化TaobaoqimenitemstorequeryAPIRequest对象
func NewTaobaoqimenitemstorequeryRequest() *TaobaoqimenitemstorequeryAPIRequest {
	return &TaobaoqimenitemstorequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenitemstorequeryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.itemstore.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenitemstorequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenitemstorequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPage is Page Setter
// 当前查询的页面编码
func (r *TaobaoqimenitemstorequeryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r TaobaoqimenitemstorequeryAPIRequest) GetPage() int64 {
	return r._page
}

// SetItemId is ItemId Setter
// 线上商品ID
func (r *TaobaoqimenitemstorequeryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoqimenitemstorequeryAPIRequest) GetItemId() int64 {
	return r._itemId
}
