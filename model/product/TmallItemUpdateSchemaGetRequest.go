package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫编辑商品规则获取 API请求
tmall.item.update.schema.get

Schema方式编辑天猫商品时，编辑商品规则获取
*/
type TmallItemUpdateSchemaGetRequest struct {
    model.Params
    // 需要编辑的商品ID
    _itemId   int64
    // 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。
    _categoryId   int64
    // 商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。
    _productId   int64
}

// 初始化TmallItemUpdateSchemaGetRequest对象
func NewTmallItemUpdateSchemaGetRequest() *TmallItemUpdateSchemaGetRequest{
    return &TmallItemUpdateSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemUpdateSchemaGetRequest) GetApiMethodName() string {
    return "tmall.item.update.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemUpdateSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 需要编辑的商品ID
func (r *TmallItemUpdateSchemaGetRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallItemUpdateSchemaGetRequest) GetItemId() int64 {
    return r._itemId
}
// CategoryId Setter
// 商品发布的目标类目，必须是叶子类目。如果没有切换类目需求，不需要填写。
func (r *TmallItemUpdateSchemaGetRequest) SetCategoryId(_categoryId int64) error {
    r._categoryId = _categoryId
    r.Set("category_id", _categoryId)
    return nil
}

// CategoryId Getter
func (r TmallItemUpdateSchemaGetRequest) GetCategoryId() int64 {
    return r._categoryId
}
// ProductId Setter
// 商品发布的目标product_id。如果没有切换产品的需求，参数可以不填写。
func (r *TmallItemUpdateSchemaGetRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallItemUpdateSchemaGetRequest) GetProductId() int64 {
    return r._productId
}
