package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品信息获取schema获取 APIRequest
tmall.product.schema.get

产品信息获取接口schema形式返回
*/
type TmallProductSchemaGetRequest struct {
    model.Params

    // 产品编号
    productId   int64 

}

func NewTmallProductSchemaGetRequest() *TmallProductSchemaGetRequest{
    return &TmallProductSchemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallProductSchemaGetRequest) GetApiMethodName() string {
    return "tmall.product.schema.get"
}

func (r TmallProductSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallProductSchemaGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallProductSchemaGetRequest) GetProductId() int64 {
    return r.productId
}

