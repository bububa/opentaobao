package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallsupplychainchannelproductquantityupdateAPIRequest 渠道无仓库存更新接口 API请求
// tmall.supplychain.channel.product.quantity.update
//
// 渠道无仓库存更新接口
type TmallsupplychainchannelproductquantityupdateAPIRequest struct {
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

// NewTmallsupplychainchannelproductquantityupdateRequest 初始化TmallsupplychainchannelproductquantityupdateAPIRequest对象
func NewTmallsupplychainchannelproductquantityupdateRequest() *TmallsupplychainchannelproductquantityupdateAPIRequest {
	return &TmallsupplychainchannelproductquantityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallsupplychainchannelproductquantityupdateAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.quantity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallsupplychainchannelproductquantityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallsupplychainchannelproductquantityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuantity is Quantity Setter
// 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0
func (r *TmallsupplychainchannelproductquantityupdateAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TmallsupplychainchannelproductquantityupdateAPIRequest) GetQuantity() int64 {
	return r._quantity
}

// SetProductId is ProductId Setter
// 产品数字ID
func (r *TmallsupplychainchannelproductquantityupdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallsupplychainchannelproductquantityupdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetSkuId is SkuId Setter
// 产品SKU ID
func (r *TmallsupplychainchannelproductquantityupdateAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TmallsupplychainchannelproductquantityupdateAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetUpdateType is UpdateType Setter
// 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新
func (r *TmallsupplychainchannelproductquantityupdateAPIRequest) SetUpdateType(_updateType int64) error {
	r._updateType = _updateType
	r.Set("update_type", _updateType)
	return nil
}

// GetUpdateType UpdateType Getter
func (r TmallsupplychainchannelproductquantityupdateAPIRequest) GetUpdateType() int64 {
	return r._updateType
}
