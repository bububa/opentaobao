package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品更新规则获取接口 APIRequest
tmall.product.update.schema.get

获取用户更新产品的规则
*/
type TmallProductUpdateSchemaGetRequest struct {
    model.Params

    // 产品编号
    productId   int64 

}

func NewTmallProductUpdateSchemaGetRequest() *TmallProductUpdateSchemaGetRequest{
    return &TmallProductUpdateSchemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallProductUpdateSchemaGetRequest) GetApiMethodName() string {
    return "tmall.product.update.schema.get"
}

func (r TmallProductUpdateSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallProductUpdateSchemaGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallProductUpdateSchemaGetRequest) GetProductId() int64 {
    return r.productId
}

