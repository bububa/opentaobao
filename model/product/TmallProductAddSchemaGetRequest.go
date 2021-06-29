package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品发布规则获取接口 API请求
tmall.product.add.schema.get

获取用户发布产品的规则
*/
type TmallProductAddSchemaGetRequest struct {
    model.Params
    // 商品发布的目标类目，必须是叶子类目
    _categoryId   int64
    // 品牌ID
    _brandId   int64
}

// 初始化TmallProductAddSchemaGetRequest对象
func NewTmallProductAddSchemaGetRequest() *TmallProductAddSchemaGetRequest{
    return &TmallProductAddSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductAddSchemaGetRequest) GetApiMethodName() string {
    return "tmall.product.add.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductAddSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallProductAddSchemaGetRequest) SetCategoryId(_categoryId int64) error {
    r._categoryId = _categoryId
    r.Set("category_id", _categoryId)
    return nil
}

// CategoryId Getter
func (r TmallProductAddSchemaGetRequest) GetCategoryId() int64 {
    return r._categoryId
}
// BrandId Setter
// 品牌ID
func (r *TmallProductAddSchemaGetRequest) SetBrandId(_brandId int64) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r TmallProductAddSchemaGetRequest) GetBrandId() int64 {
    return r._brandId
}
