package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫编辑商品规则获取 APIRequest
tmall.item.update.schema.get

Schema方式编辑天猫商品时，编辑商品规则获取
*/
type TmallItemUpdateSchemaGetRequest struct {
    model.Params

    // 需要编辑的商品ID
    itemId   int64 

    // 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。
    categoryId   int64 

    // 商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。
    productId   int64 

}

func NewTmallItemUpdateSchemaGetRequest() *TmallItemUpdateSchemaGetRequest{
    return &TmallItemUpdateSchemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemUpdateSchemaGetRequest) GetApiMethodName() string {
    return "tmall.item.update.schema.get"
}

func (r TmallItemUpdateSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemUpdateSchemaGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallItemUpdateSchemaGetRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TmallItemUpdateSchemaGetRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r TmallItemUpdateSchemaGetRequest) GetCategoryId() int64 {
    return r.categoryId
}

func (r *TmallItemUpdateSchemaGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallItemUpdateSchemaGetRequest) GetProductId() int64 {
    return r.productId
}

