package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品SKU删除接口 APIRequest
taobao.fenxiao.product.sku.delete

根据sku properties删除sku数据
*/
type TaobaoFenxiaoProductSkuDeleteRequest struct {
    model.Params

    // 产品id
    productId   int64 

    // sku属性
    properties   string 

}

func NewTaobaoFenxiaoProductSkuDeleteRequest() *TaobaoFenxiaoProductSkuDeleteRequest{
    return &TaobaoFenxiaoProductSkuDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoProductSkuDeleteRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.sku.delete"
}

func (r TaobaoFenxiaoProductSkuDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoProductSkuDeleteRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoFenxiaoProductSkuDeleteRequest) GetProductId() int64 {
    return r.productId
}

func (r *TaobaoFenxiaoProductSkuDeleteRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

func (r TaobaoFenxiaoProductSkuDeleteRequest) GetProperties() string {
    return r.properties
}

