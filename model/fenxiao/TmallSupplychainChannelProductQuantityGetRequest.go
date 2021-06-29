package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道库存查询接口 API请求
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

// 初始化TmallSupplychainChannelProductQuantityGetRequest对象
func NewTmallSupplychainChannelProductQuantityGetRequest() *TmallSupplychainChannelProductQuantityGetRequest{
    return &TmallSupplychainChannelProductQuantityGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductQuantityGetRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.quantity.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductQuantityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品数字ID
func (r *TmallSupplychainChannelProductQuantityGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TmallSupplychainChannelProductQuantityGetRequest) GetProductId() int64 {
    return r.productId
}
// SkuId Setter
// SKU ID
func (r *TmallSupplychainChannelProductQuantityGetRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

// SkuId Getter
func (r TmallSupplychainChannelProductQuantityGetRequest) GetSkuId() int64 {
    return r.skuId
}
