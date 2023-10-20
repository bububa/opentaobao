package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSchemaAddAPIRequest 天猫根据规则发布商品 API请求
// tmall.item.schema.add
//
// 天猫TopSchema发布商品。
type TmallItemSchemaAddAPIRequest struct {
	model.Params
	// 根据tmall.item.add.schema.get生成的商品发布规则入参数据
	_xmlData string
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
	// 发布商品的productId，如果tmall.product.match.schema.get获取到得字段为空，这个参数传入0，否则需要通过tmall.product.schema.match查询到得可用productId
	_productId int64
}

// NewTmallItemSchemaAddRequest 初始化TmallItemSchemaAddAPIRequest对象
func NewTmallItemSchemaAddRequest() *TmallItemSchemaAddAPIRequest {
	return &TmallItemSchemaAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSchemaAddAPIRequest) GetApiMethodName() string {
	return "tmall.item.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSchemaAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSchemaAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXmlData is XmlData Setter
// 根据tmall.item.add.schema.get生成的商品发布规则入参数据
func (r *TmallItemSchemaAddAPIRequest) SetXmlData(_xmlData string) error {
	r._xmlData = _xmlData
	r.Set("xml_data", _xmlData)
	return nil
}

// GetXmlData XmlData Getter
func (r TmallItemSchemaAddAPIRequest) GetXmlData() string {
	return r._xmlData
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallItemSchemaAddAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallItemSchemaAddAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetProductId is ProductId Setter
// 发布商品的productId，如果tmall.product.match.schema.get获取到得字段为空，这个参数传入0，否则需要通过tmall.product.schema.match查询到得可用productId
func (r *TmallItemSchemaAddAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallItemSchemaAddAPIRequest) GetProductId() int64 {
	return r._productId
}
