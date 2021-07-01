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
type TmallProductUpdateSchemaGetAPIRequest struct {
    model.Params
    // 产品编号
    _productId   int64
}

// 初始化TmallProductUpdateSchemaGetAPIRequest对象
func NewTmallProductUpdateSchemaGetRequest() *TmallProductUpdateSchemaGetAPIRequest{
    return &TmallProductUpdateSchemaGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductUpdateSchemaGetAPIRequest) GetApiMethodName() string {
    return "tmall.product.update.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductUpdateSchemaGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品编号
func (r *TmallProductUpdateSchemaGetAPIRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallProductUpdateSchemaGetAPIRequest) GetProductId() int64 {
    return r._productId
}
