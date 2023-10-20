package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallsupplychainchannelproductquantitygetAPIRequest 渠道库存查询接口 API请求
// tmall.supplychain.channel.product.quantity.get
//
// 渠道库存查询接口
type TmallsupplychainchannelproductquantitygetAPIRequest struct {
	model.Params
	// 产品数字ID
	_productId int64
	// SKU ID
	_skuId int64
}

// NewTmallsupplychainchannelproductquantitygetRequest 初始化TmallsupplychainchannelproductquantitygetAPIRequest对象
func NewTmallsupplychainchannelproductquantitygetRequest() *TmallsupplychainchannelproductquantitygetAPIRequest {
	return &TmallsupplychainchannelproductquantitygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallsupplychainchannelproductquantitygetAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.quantity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallsupplychainchannelproductquantitygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallsupplychainchannelproductquantitygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品数字ID
func (r *TmallsupplychainchannelproductquantitygetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallsupplychainchannelproductquantitygetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetSkuId is SkuId Setter
// SKU ID
func (r *TmallsupplychainchannelproductquantitygetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TmallsupplychainchannelproductquantitygetAPIRequest) GetSkuId() int64 {
	return r._skuId
}
