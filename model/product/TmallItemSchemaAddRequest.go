package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫根据规则发布商品 API请求
tmall.item.schema.add

天猫TopSchema发布商品。
*/
type TmallItemSchemaAddRequest struct {
    model.Params
    // 商品发布的目标类目，必须是叶子类目
    categoryId   int64
    // 发布商品的productId，如果tmall.product.match.schema.get获取到得字段为空，这个参数传入0，否则需要通过tmall.product.schema.match查询到得可用productId
    productId   int64
    // 根据tmall.item.add.schema.get生成的商品发布规则入参数据
    xmlData   string
}

// 初始化TmallItemSchemaAddRequest对象
func NewTmallItemSchemaAddRequest() *TmallItemSchemaAddRequest{
    return &TmallItemSchemaAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSchemaAddRequest) GetApiMethodName() string {
    return "tmall.item.schema.add"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSchemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallItemSchemaAddRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r TmallItemSchemaAddRequest) GetCategoryId() int64 {
    return r.categoryId
}
// ProductId Setter
// 发布商品的productId，如果tmall.product.match.schema.get获取到得字段为空，这个参数传入0，否则需要通过tmall.product.schema.match查询到得可用productId
func (r *TmallItemSchemaAddRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TmallItemSchemaAddRequest) GetProductId() int64 {
    return r.productId
}
// XmlData Setter
// 根据tmall.item.add.schema.get生成的商品发布规则入参数据
func (r *TmallItemSchemaAddRequest) SetXmlData(xmlData string) error {
    r.xmlData = xmlData
    r.Set("xml_data", xmlData)
    return nil
}

// XmlData Getter
func (r TmallItemSchemaAddRequest) GetXmlData() string {
    return r.xmlData
}
