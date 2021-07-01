package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
使用Schema文件发布一个产品 API请求
tmall.product.schema.add

Schema体系发布一个产品
*/
type TmallProductSchemaAddAPIRequest struct {
    model.Params
    // 商品发布的目标类目，必须是叶子类目
    _categoryId   int64
    // 品牌ID
    _brandId   int64
    // 根据tmall.product.add.schema.get生成的产品发布规则入参数据
    _xmlData   string
}

// 初始化TmallProductSchemaAddAPIRequest对象
func NewTmallProductSchemaAddRequest() *TmallProductSchemaAddAPIRequest{
    return &TmallProductSchemaAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductSchemaAddAPIRequest) GetApiMethodName() string {
    return "tmall.product.schema.add"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductSchemaAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallProductSchemaAddAPIRequest) SetCategoryId(_categoryId int64) error {
    r._categoryId = _categoryId
    r.Set("category_id", _categoryId)
    return nil
}

// CategoryId Getter
func (r TmallProductSchemaAddAPIRequest) GetCategoryId() int64 {
    return r._categoryId
}
// BrandId Setter
// 品牌ID
func (r *TmallProductSchemaAddAPIRequest) SetBrandId(_brandId int64) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r TmallProductSchemaAddAPIRequest) GetBrandId() int64 {
    return r._brandId
}
// XmlData Setter
// 根据tmall.product.add.schema.get生成的产品发布规则入参数据
func (r *TmallProductSchemaAddAPIRequest) SetXmlData(_xmlData string) error {
    r._xmlData = _xmlData
    r.Set("xml_data", _xmlData)
    return nil
}

// XmlData Getter
func (r TmallProductSchemaAddAPIRequest) GetXmlData() string {
    return r._xmlData
}
