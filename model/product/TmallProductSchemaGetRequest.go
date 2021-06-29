package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品信息获取schema获取 API请求
tmall.product.schema.get

产品信息获取接口schema形式返回
*/
type TmallProductSchemaGetRequest struct {
    model.Params
    // 产品编号
    productId   int64
}

// 初始化TmallProductSchemaGetRequest对象
func NewTmallProductSchemaGetRequest() *TmallProductSchemaGetRequest{
    return &TmallProductSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductSchemaGetRequest) GetApiMethodName() string {
    return "tmall.product.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品编号
func (r *TmallProductSchemaGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TmallProductSchemaGetRequest) GetProductId() int64 {
    return r.productId
}
