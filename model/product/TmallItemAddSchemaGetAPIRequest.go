package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemaddschemagetAPIRequest 天猫发布商品规则获取 API请求
// tmall.item.add.schema.get
//
// 通过类目以及productId获取商品发布规则；
type TmallitemaddschemagetAPIRequest struct {
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

// NewTmallitemaddschemagetRequest 初始化TmallitemaddschemagetAPIRequest对象
func NewTmallitemaddschemagetRequest() *TmallitemaddschemagetAPIRequest {
	return &TmallitemaddschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemaddschemagetAPIRequest) GetApiMethodName() string {
	return "tmall.item.add.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemaddschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemaddschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 发布商品类型，一口价填“b”，拍卖填&#34;a&#34;
func (r *TmallitemaddschemagetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TmallitemaddschemagetAPIRequest) GetType() string {
	return r._type
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallitemaddschemagetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallitemaddschemagetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetProductId is ProductId Setter
// 商品发布的目标product_id
func (r *TmallitemaddschemagetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallitemaddschemagetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetIsvInit is IsvInit Setter
// 正常接口调用时，请忽略这个参数或者填FALSE。这个参数提供给ISV对接Schema时，如果想先获取了解所有字段和规则，可以将此字段设置为true，product_id也就不需要提供了，设置为0即可
func (r *TmallitemaddschemagetAPIRequest) SetIsvInit(_isvInit bool) error {
	r._isvInit = _isvInit
	r.Set("isv_init", _isvInit)
	return nil
}

// GetIsvInit IsvInit Getter
func (r TmallitemaddschemagetAPIRequest) GetIsvInit() bool {
	return r._isvInit
}
