package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫根据规则发布商品 APIRequest
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

func NewTmallItemSchemaAddRequest() *TmallItemSchemaAddRequest{
    return &TmallItemSchemaAddRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemSchemaAddRequest) GetApiMethodName() string {
    return "tmall.item.schema.add"
}

func (r TmallItemSchemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemSchemaAddRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r TmallItemSchemaAddRequest) GetCategoryId() int64 {
    return r.categoryId
}

func (r *TmallItemSchemaAddRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallItemSchemaAddRequest) GetProductId() int64 {
    return r.productId
}

func (r *TmallItemSchemaAddRequest) SetXmlData(xmlData string) error {
    r.xmlData = xmlData
    r.Set("xml_data", xmlData)
    return nil
}

func (r TmallItemSchemaAddRequest) GetXmlData() string {
    return r.xmlData
}

