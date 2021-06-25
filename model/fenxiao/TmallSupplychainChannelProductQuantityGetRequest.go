package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道库存查询接口 APIRequest
tmall.supplychain.channel.product.quantity.get

渠道库存查询接口
*/
type TmallSupplychainChannelProductQuantityGetRequest struct {
    model.Params

    // 产品数字ID
    productId   int64 

    // SKU ID
    skuId   int64 

}

func NewTmallSupplychainChannelProductQuantityGetRequest() *TmallSupplychainChannelProductQuantityGetRequest{
    return &TmallSupplychainChannelProductQuantityGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallSupplychainChannelProductQuantityGetRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.quantity.get"
}

func (r TmallSupplychainChannelProductQuantityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallSupplychainChannelProductQuantityGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallSupplychainChannelProductQuantityGetRequest) GetProductId() int64 {
    return r.productId
}

func (r *TmallSupplychainChannelProductQuantityGetRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

func (r TmallSupplychainChannelProductQuantityGetRequest) GetSkuId() int64 {
    return r.skuId
}

