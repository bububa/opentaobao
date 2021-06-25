package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫根据规则编辑商品 APIRequest
tmall.item.schema.update

天猫根据规则编辑商品
*/
type TmallItemSchemaUpdateRequest struct {
    model.Params

    // 需要编辑的商品ID
    itemId   int64 

    // 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求不需要填写
    categoryId   int64 

    // 商品发布的目标product_id。如果没有切换类目或者切换产品的需求，参数不用填写
    productId   int64 

    // 根据tmall.item.update.schema.get生成的商品编辑规则入参数据
    xmlData   string 

}

func NewTmallItemSchemaUpdateRequest() *TmallItemSchemaUpdateRequest{
    return &TmallItemSchemaUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemSchemaUpdateRequest) GetApiMethodName() string {
    return "tmall.item.schema.update"
}

func (r TmallItemSchemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemSchemaUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallItemSchemaUpdateRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TmallItemSchemaUpdateRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r TmallItemSchemaUpdateRequest) GetCategoryId() int64 {
    return r.categoryId
}

func (r *TmallItemSchemaUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallItemSchemaUpdateRequest) GetProductId() int64 {
    return r.productId
}

func (r *TmallItemSchemaUpdateRequest) SetXmlData(xmlData string) error {
    r.xmlData = xmlData
    r.Set("xml_data", xmlData)
    return nil
}

func (r TmallItemSchemaUpdateRequest) GetXmlData() string {
    return r.xmlData
}

