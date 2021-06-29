package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫发布商品规则获取 API请求
tmall.item.add.schema.get

通过类目以及productId获取商品发布规则；
*/
type TmallItemAddSchemaGetRequest struct {
    model.Params
    // 商品发布的目标类目，必须是叶子类目
    categoryId   int64
    // 商品发布的目标product_id
    productId   int64
    // 发布商品类型，一口价填“b”，拍卖填"a"
    type   string
    // 正常接口调用时，请忽略这个参数或者填FALSE。这个参数提供给ISV对接Schema时，如果想先获取了解所有字段和规则，可以将此字段设置为true，product_id也就不需要提供了，设置为0即可
    isvInit   bool
}

// 初始化TmallItemAddSchemaGetRequest对象
func NewTmallItemAddSchemaGetRequest() *TmallItemAddSchemaGetRequest{
    return &TmallItemAddSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemAddSchemaGetRequest) GetApiMethodName() string {
    return "tmall.item.add.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemAddSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallItemAddSchemaGetRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r TmallItemAddSchemaGetRequest) GetCategoryId() int64 {
    return r.categoryId
}
// ProductId Setter
// 商品发布的目标product_id
func (r *TmallItemAddSchemaGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TmallItemAddSchemaGetRequest) GetProductId() int64 {
    return r.productId
}
// Type Setter
// 发布商品类型，一口价填“b”，拍卖填"a"
func (r *TmallItemAddSchemaGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TmallItemAddSchemaGetRequest) GetType() string {
    return r.type
}
// IsvInit Setter
// 正常接口调用时，请忽略这个参数或者填FALSE。这个参数提供给ISV对接Schema时，如果想先获取了解所有字段和规则，可以将此字段设置为true，product_id也就不需要提供了，设置为0即可
func (r *TmallItemAddSchemaGetRequest) SetIsvInit(isvInit bool) error {
    r.isvInit = isvInit
    r.Set("isv_init", isvInit)
    return nil
}

// IsvInit Getter
func (r TmallItemAddSchemaGetRequest) GetIsvInit() bool {
    return r.isvInit
}
