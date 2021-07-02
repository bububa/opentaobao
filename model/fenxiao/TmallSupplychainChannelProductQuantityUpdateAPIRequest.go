package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductQuantityUpdateAPIRequest 渠道无仓库存更新接口 API请求
// tmall.supplychain.channel.product.quantity.update
//
// 渠道无仓库存更新接口
type TmallSupplychainChannelProductQuantityUpdateAPIRequest struct {
	model.Params
	// 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
	_quantity int64
	// 产品数字ID
	_productId int64
	// 产品SKU ID
	_skuId int64
	// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
	_updateType int64
}

// NewTmallSupplychainChannelProductQuantityUpdateRequest 初始化TmallSupplychainChannelProductQuantityUpdateAPIRequest对象
func NewTmallSupplychainChannelProductQuantityUpdateRequest() *TmallSupplychainChannelProductQuantityUpdateAPIRequest {
	return &TmallSupplychainChannelProductQuantityUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductQuantityUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.quantity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductQuantityUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Quantity Setter
// 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
func (r *TmallSupplychainChannelProductQuantityUpdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// Get Quantity Getter
func (r TmallSupplychainChannelProductQuantityUpdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}

// Set is ProductId Setter
// 产品数字ID
func (r *TmallSupplychainChannelProductQuantityUpdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r TmallSupplychainChannelProductQuantityUpdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// Set is SkuId Setter
// 产品SKU ID
func (r *TmallSupplychainChannelProductQuantityUpdateAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// Get SkuId Getter
func (r TmallSupplychainChannelProductQuantityUpdateAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// Set is UpdateType Setter
// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
func (r *TmallSupplychainChannelProductQuantityUpdateAPIRequest) SetUpdateType(_updateType int64) error {
	r._updateType = _updateType
	r.Set("update_type", _updateType)
	return nil
}

// Get UpdateType Getter
func (r TmallSupplychainChannelProductQuantityUpdateAPIRequest) GetUpdateType() int64 {
	return r._updateType
}
