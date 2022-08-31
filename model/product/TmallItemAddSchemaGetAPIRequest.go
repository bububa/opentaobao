package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemAddSchemaGetAPIRequest 天猫发布商品规则获取 API请求
// tmall.item.add.schema.get
//
// 通过类目以及productId获取商品发布规则；
type TmallItemAddSchemaGetAPIRequest struct {
	model.Params
	// 发布商品类型，一口价填“b”，拍卖填"a"
	_type string
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
	// 商品发布的目标product_id
	_productId int64
	// 正常接口调用时，请忽略这个参数或者填FALSE。这个参数提供给ISV对接Schema时，如果想先获取了解所有字段和规则，可以将此字段设置为true，product_id也就不需要提供了，设置为0即可
	_isvInit bool
}

// NewTmallItemAddSchemaGetRequest 初始化TmallItemAddSchemaGetAPIRequest对象
func NewTmallItemAddSchemaGetRequest() *TmallItemAddSchemaGetAPIRequest {
	return &TmallItemAddSchemaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemAddSchemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.add.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemAddSchemaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetType is Type Setter
// 发布商品类型，一口价填“b”，拍卖填&#34;a&#34;
func (r *TmallItemAddSchemaGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TmallItemAddSchemaGetAPIRequest) GetType() string {
	return r._type
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallItemAddSchemaGetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallItemAddSchemaGetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetProductId is ProductId Setter
// 商品发布的目标product_id
func (r *TmallItemAddSchemaGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallItemAddSchemaGetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetIsvInit is IsvInit Setter
// 正常接口调用时，请忽略这个参数或者填FALSE。这个参数提供给ISV对接Schema时，如果想先获取了解所有字段和规则，可以将此字段设置为true，product_id也就不需要提供了，设置为0即可
func (r *TmallItemAddSchemaGetAPIRequest) SetIsvInit(_isvInit bool) error {
	r._isvInit = _isvInit
	r.Set("isv_init", _isvInit)
	return nil
}

// GetIsvInit IsvInit Getter
func (r TmallItemAddSchemaGetAPIRequest) GetIsvInit() bool {
	return r._isvInit
}
