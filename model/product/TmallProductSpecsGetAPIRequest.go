package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallproductspecsgetAPIRequest 获取产品的规格信息 API请求
// tmall.product.specs.get
//
// 按product_id或品牌下载产品规格，返回一组的产品规格信息。
type TmallproductspecsgetAPIRequest struct {
	model.Params
	// 关键属性的字符串，pid:vid;pid:vid该字段必须和cat_id同时传入才起效果。 而且只有当product_id不传入的时候才起效果。<br>product_id 不传时该参数必传
	_properties string
	// 产品的ID。这个不能和properties和cat_id同时起效果<br>properties 和cat_id 均不传时，该参数必传。
	_productId int64
	// 类目的ID号，该id必须和properties同时传入。而且只有当product_id不传入的时候才起效果。<br> product_id不传时，该参数必传
	_catId int64
}

// NewTmallproductspecsgetRequest 初始化TmallproductspecsgetAPIRequest对象
func NewTmallproductspecsgetRequest() *TmallproductspecsgetAPIRequest {
	return &TmallproductspecsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallproductspecsgetAPIRequest) GetApiMethodName() string {
	return "tmall.product.specs.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallproductspecsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallproductspecsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProperties is Properties Setter
// 关键属性的字符串，pid:vid;pid:vid该字段必须和cat_id同时传入才起效果。 而且只有当product_id不传入的时候才起效果。&lt;br&gt;product_id 不传时该参数必传
func (r *TmallproductspecsgetAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TmallproductspecsgetAPIRequest) GetProperties() string {
	return r._properties
}

// SetProductId is ProductId Setter
// 产品的ID。这个不能和properties和cat_id同时起效果&lt;br&gt;properties 和cat_id 均不传时，该参数必传。
func (r *TmallproductspecsgetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallproductspecsgetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetCatId is CatId Setter
// 类目的ID号，该id必须和properties同时传入。而且只有当product_id不传入的时候才起效果。&lt;br&gt; product_id不传时，该参数必传
func (r *TmallproductspecsgetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TmallproductspecsgetAPIRequest) GetCatId() int64 {
	return r._catId
}
