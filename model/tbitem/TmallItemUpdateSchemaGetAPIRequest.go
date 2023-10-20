package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemupdateschemagetAPIRequest 天猫编辑商品规则获取 API请求
// tmall.item.update.schema.get
//
// Schema方式编辑天猫商品时，编辑商品规则获取
type TmallitemupdateschemagetAPIRequest struct {
	model.Params
	// 需要编辑的商品ID
	_itemId int64
	// 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。
	_categoryId int64
	// 商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。
	_productId int64
}

// NewTmallitemupdateschemagetRequest 初始化TmallitemupdateschemagetAPIRequest对象
func NewTmallitemupdateschemagetRequest() *TmallitemupdateschemagetAPIRequest {
	return &TmallitemupdateschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemupdateschemagetAPIRequest) GetApiMethodName() string {
	return "tmall.item.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemupdateschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemupdateschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 需要编辑的商品ID
func (r *TmallitemupdateschemagetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemupdateschemagetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。
func (r *TmallitemupdateschemagetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallitemupdateschemagetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetProductId is ProductId Setter
// 商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。
func (r *TmallitemupdateschemagetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallitemupdateschemagetAPIRequest) GetProductId() int64 {
	return r._productId
}
