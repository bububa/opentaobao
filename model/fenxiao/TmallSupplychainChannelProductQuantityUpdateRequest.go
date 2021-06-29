package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道无仓库存更新接口 API请求
tmall.supplychain.channel.product.quantity.update

渠道无仓库存更新接口
*/
type TmallSupplychainChannelProductQuantityUpdateRequest struct {
    model.Params
    // 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
    _quantity   int64
    // 产品数字ID
    _productId   int64
    // 产品SKU ID
    _skuId   int64
    // 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
    _updateType   int64
}

// 初始化TmallSupplychainChannelProductQuantityUpdateRequest对象
func NewTmallSupplychainChannelProductQuantityUpdateRequest() *TmallSupplychainChannelProductQuantityUpdateRequest{
    return &TmallSupplychainChannelProductQuantityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductQuantityUpdateRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.quantity.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductQuantityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Quantity Setter
// 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
func (r *TmallSupplychainChannelProductQuantityUpdateRequest) SetQuantity(_quantity int64) error {
    r._quantity = _quantity
    r.Set("quantity", _quantity)
    return nil
}

// Quantity Getter
func (r TmallSupplychainChannelProductQuantityUpdateRequest) GetQuantity() int64 {
    return r._quantity
}
// ProductId Setter
// 产品数字ID
func (r *TmallSupplychainChannelProductQuantityUpdateRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallSupplychainChannelProductQuantityUpdateRequest) GetProductId() int64 {
    return r._productId
}
// SkuId Setter
// 产品SKU ID
func (r *TmallSupplychainChannelProductQuantityUpdateRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TmallSupplychainChannelProductQuantityUpdateRequest) GetSkuId() int64 {
    return r._skuId
}
// UpdateType Setter
// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
func (r *TmallSupplychainChannelProductQuantityUpdateRequest) SetUpdateType(_updateType int64) error {
    r._updateType = _updateType
    r.Set("update_type", _updateType)
    return nil
}

// UpdateType Getter
func (r TmallSupplychainChannelProductQuantityUpdateRequest) GetUpdateType() int64 {
    return r._updateType
}
