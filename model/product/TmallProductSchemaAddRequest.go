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
type TmallProductSchemaAddRequest struct {
    model.Params
    // 商品发布的目标类目，必须是叶子类目
    categoryId   int64
    // 品牌ID
    brandId   int64
    // 根据tmall.product.add.schema.get生成的产品发布规则入参数据
    xmlData   string
}

// 初始化TmallProductSchemaAddRequest对象
func NewTmallProductSchemaAddRequest() *TmallProductSchemaAddRequest{
    return &TmallProductSchemaAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductSchemaAddRequest) GetApiMethodName() string {
    return "tmall.product.schema.add"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductSchemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallProductSchemaAddRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r TmallProductSchemaAddRequest) GetCategoryId() int64 {
    return r.categoryId
}
// BrandId Setter
// 品牌ID
func (r *TmallProductSchemaAddRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

// BrandId Getter
func (r TmallProductSchemaAddRequest) GetBrandId() int64 {
    return r.brandId
}
// XmlData Setter
// 根据tmall.product.add.schema.get生成的产品发布规则入参数据
func (r *TmallProductSchemaAddRequest) SetXmlData(xmlData string) error {
    r.xmlData = xmlData
    r.Set("xml_data", xmlData)
    return nil
}

// XmlData Getter
func (r TmallProductSchemaAddRequest) GetXmlData() string {
    return r.xmlData
}
