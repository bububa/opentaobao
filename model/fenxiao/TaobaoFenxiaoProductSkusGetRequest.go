package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
SKU查询接口 APIRequest
taobao.fenxiao.product.skus.get

产品sku查询
*/
type TaobaoFenxiaoProductSkusGetRequest struct {
    model.Params

    // 产品ID
    productId   int64 

}

func NewTaobaoFenxiaoProductSkusGetRequest() *TaobaoFenxiaoProductSkusGetRequest{
    return &TaobaoFenxiaoProductSkusGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoProductSkusGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.skus.get"
}

func (r TaobaoFenxiaoProductSkusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoProductSkusGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoFenxiaoProductSkusGetRequest) GetProductId() int64 {
    return r.productId
}

