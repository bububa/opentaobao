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
type TmallSupplychainChannelProductQuantityGetAPIRequest struct {
    model.Params
    // 产品数字ID
    _productId   int64
    // SKU ID
    _skuId   int64
}

// 初始化TmallSupplychainChannelProductQuantityGetAPIRequest对象
func NewTmallSupplychainChannelProductQuantityGetRequest() *TmallSupplychainChannelProductQuantityGetAPIRequest{
    return &TmallSupplychainChannelProductQuantityGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductQuantityGetAPIRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.quantity.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductQuantityGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品数字ID
func (r *TmallSupplychainChannelProductQuantityGetAPIRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallSupplychainChannelProductQuantityGetAPIRequest) GetProductId() int64 {
    return r._productId
}
// SkuId Setter
// SKU ID
func (r *TmallSupplychainChannelProductQuantityGetAPIRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TmallSupplychainChannelProductQuantityGetAPIRequest) GetSkuId() int64 {
    return r._skuId
}
