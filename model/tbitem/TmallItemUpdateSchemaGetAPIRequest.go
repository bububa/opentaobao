package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemUpdateSchemaGetAPIRequest 天猫编辑商品规则获取 API请求
// tmall.item.update.schema.get
//
// Schema方式编辑天猫商品时，编辑商品规则获取
type TmallItemUpdateSchemaGetAPIRequest struct {
	model.Params
	// 需要编辑的商品ID
	_itemId int64
	// 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。
	_categoryId int64
	// 商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。
	_productId int64
}

// NewTmallItemUpdateSchemaGetRequest 初始化TmallItemUpdateSchemaGetAPIRequest对象
func NewTmallItemUpdateSchemaGetRequest() *TmallItemUpdateSchemaGetAPIRequest {
	return &TmallItemUpdateSchemaGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemUpdateSchemaGetAPIRequest) Reset() {
	r._itemId = 0
	r._categoryId = 0
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemUpdateSchemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemUpdateSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemUpdateSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 需要编辑的商品ID
func (r *TmallItemUpdateSchemaGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemUpdateSchemaGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。
func (r *TmallItemUpdateSchemaGetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallItemUpdateSchemaGetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetProductId is ProductId Setter
// 商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。
func (r *TmallItemUpdateSchemaGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallItemUpdateSchemaGetAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolTmallItemUpdateSchemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemUpdateSchemaGetRequest()
	},
}

// GetTmallItemUpdateSchemaGetRequest 从 sync.Pool 获取 TmallItemUpdateSchemaGetAPIRequest
func GetTmallItemUpdateSchemaGetAPIRequest() *TmallItemUpdateSchemaGetAPIRequest {
	return poolTmallItemUpdateSchemaGetAPIRequest.Get().(*TmallItemUpdateSchemaGetAPIRequest)
}

// ReleaseTmallItemUpdateSchemaGetAPIRequest 将 TmallItemUpdateSchemaGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemUpdateSchemaGetAPIRequest(v *TmallItemUpdateSchemaGetAPIRequest) {
	v.Reset()
	poolTmallItemUpdateSchemaGetAPIRequest.Put(v)
}
