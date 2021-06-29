package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品更新规则获取接口 API请求
tmall.product.update.schema.get

获取用户更新产品的规则
*/
type TmallProductUpdateSchemaGetRequest struct {
    model.Params
    // 产品编号
    productId   int64
}

// 初始化TmallProductUpdateSchemaGetRequest对象
func NewTmallProductUpdateSchemaGetRequest() *TmallProductUpdateSchemaGetRequest{
    return &TmallProductUpdateSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductUpdateSchemaGetRequest) GetApiMethodName() string {
    return "tmall.product.update.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductUpdateSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品编号
func (r *TmallProductUpdateSchemaGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TmallProductUpdateSchemaGetRequest) GetProductId() int64 {
    return r.productId
}
