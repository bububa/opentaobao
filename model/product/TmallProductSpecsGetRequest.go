package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品的规格信息 API请求
tmall.product.specs.get

按product_id或品牌下载产品规格，返回一组的产品规格信息。
*/
type TmallProductSpecsGetRequest struct {
    model.Params
    // 产品的ID。这个不能和properties和cat_id同时起效果<br>properties 和cat_id 均不传时，该参数必传。
    _productId   int64
    // 关键属性的字符串，pid:vid;pid:vid该字段必须和cat_id同时传入才起效果。 而且只有当product_id不传入的时候才起效果。<br>product_id 不传时该参数必传
    _properties   string
    // 类目的ID号，该id必须和properties同时传入。而且只有当product_id不传入的时候才起效果。<br> product_id不传时，该参数必传
    _catId   int64
}

// 初始化TmallProductSpecsGetRequest对象
func NewTmallProductSpecsGetRequest() *TmallProductSpecsGetRequest{
    return &TmallProductSpecsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductSpecsGetRequest) GetApiMethodName() string {
    return "tmall.product.specs.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductSpecsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品的ID。这个不能和properties和cat_id同时起效果<br>properties 和cat_id 均不传时，该参数必传。
func (r *TmallProductSpecsGetRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallProductSpecsGetRequest) GetProductId() int64 {
    return r._productId
}
// Properties Setter
// 关键属性的字符串，pid:vid;pid:vid该字段必须和cat_id同时传入才起效果。 而且只有当product_id不传入的时候才起效果。<br>product_id 不传时该参数必传
func (r *TmallProductSpecsGetRequest) SetProperties(_properties string) error {
    r._properties = _properties
    r.Set("properties", _properties)
    return nil
}

// Properties Getter
func (r TmallProductSpecsGetRequest) GetProperties() string {
    return r._properties
}
// CatId Setter
// 类目的ID号，该id必须和properties同时传入。而且只有当product_id不传入的时候才起效果。<br> product_id不传时，该参数必传
func (r *TmallProductSpecsGetRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r TmallProductSpecsGetRequest) GetCatId() int64 {
    return r._catId
}
