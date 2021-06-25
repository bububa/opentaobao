package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品的规格信息 APIRequest
tmall.product.specs.get

按product_id或品牌下载产品规格，返回一组的产品规格信息。
*/
type TmallProductSpecsGetRequest struct {
    model.Params

    // 产品的ID。这个不能和properties和cat_id同时起效果<br>
properties 和cat_id 均不传时，该参数必传。
    productId   int64 

    // 关键属性的字符串，pid:vid;pid:vid
该字段必须和cat_id同时传入才起效果。 而且只有当product_id不传入的时候才起效果。<br>product_id 不传时该参数必传
    properties   string 

    // 类目的ID号，该id必须和properties同时传入。
而且只有当product_id不传入的时候才起效果。<br> product_id不传时，该参数必传
    catId   int64 

}

func NewTmallProductSpecsGetRequest() *TmallProductSpecsGetRequest{
    return &TmallProductSpecsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallProductSpecsGetRequest) GetApiMethodName() string {
    return "tmall.product.specs.get"
}

func (r TmallProductSpecsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallProductSpecsGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallProductSpecsGetRequest) GetProductId() int64 {
    return r.productId
}

func (r *TmallProductSpecsGetRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

func (r TmallProductSpecsGetRequest) GetProperties() string {
    return r.properties
}

func (r *TmallProductSpecsGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r TmallProductSpecsGetRequest) GetCatId() int64 {
    return r.catId
}

