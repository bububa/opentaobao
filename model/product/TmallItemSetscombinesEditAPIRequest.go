package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemsetscombineseditAPIRequest 普通商品转套装商品&套装商品编辑接口 API请求
// tmall.item.setscombines.edit
//
// 普通商品转套装商品&amp;套装商品编辑接口
type TmallitemsetscombineseditAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 商品SKU更新套装货品时候用的数据
	_updateSkuScProduct *UpdateSkuScProduct
}

// NewTmallitemsetscombineseditRequest 初始化TmallitemsetscombineseditAPIRequest对象
func NewTmallitemsetscombineseditRequest() *TmallitemsetscombineseditAPIRequest {
	return &TmallitemsetscombineseditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemsetscombineseditAPIRequest) GetApiMethodName() string {
	return "tmall.item.setscombines.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemsetscombineseditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemsetscombineseditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallitemsetscombineseditAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemsetscombineseditAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetUpdateSkuScProduct is UpdateSkuScProduct Setter
// 商品SKU更新套装货品时候用的数据
func (r *TmallitemsetscombineseditAPIRequest) SetUpdateSkuScProduct(_updateSkuScProduct *UpdateSkuScProduct) error {
	r._updateSkuScProduct = _updateSkuScProduct
	r.Set("update_sku_sc_product", _updateSkuScProduct)
	return nil
}

// GetUpdateSkuScProduct UpdateSkuScProduct Getter
func (r TmallitemsetscombineseditAPIRequest) GetUpdateSkuScProduct() *UpdateSkuScProduct {
	return r._updateSkuScProduct
}
