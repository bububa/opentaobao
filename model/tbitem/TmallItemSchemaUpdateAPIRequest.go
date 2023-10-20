package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemschemaupdateAPIRequest 天猫根据规则编辑商品 API请求
// tmall.item.schema.update
//
// 天猫根据规则编辑商品
type TmallitemschemaupdateAPIRequest struct {
	model.Params
	// 根据tmall.item.update.schema.get生成的商品编辑规则入参数据
	_xmlData string
	// 需要编辑的商品ID
	_itemId int64
	// 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求不需要填写
	_categoryId int64
	// 商品发布的目标product_id。如果没有切换类目或者切换产品的需求，参数不用填写
	_productId int64
}

// NewTmallitemschemaupdateRequest 初始化TmallitemschemaupdateAPIRequest对象
func NewTmallitemschemaupdateRequest() *TmallitemschemaupdateAPIRequest {
	return &TmallitemschemaupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemschemaupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemschemaupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemschemaupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXmlData is XmlData Setter
// 根据tmall.item.update.schema.get生成的商品编辑规则入参数据
func (r *TmallitemschemaupdateAPIRequest) SetXmlData(_xmlData string) error {
	r._xmlData = _xmlData
	r.Set("xml_data", _xmlData)
	return nil
}

// GetXmlData XmlData Getter
func (r TmallitemschemaupdateAPIRequest) GetXmlData() string {
	return r._xmlData
}

// SetItemId is ItemId Setter
// 需要编辑的商品ID
func (r *TmallitemschemaupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemschemaupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求不需要填写
func (r *TmallitemschemaupdateAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallitemschemaupdateAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetProductId is ProductId Setter
// 商品发布的目标product_id。如果没有切换类目或者切换产品的需求，参数不用填写
func (r *TmallitemschemaupdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallitemschemaupdateAPIRequest) GetProductId() int64 {
	return r._productId
}
