package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDistributorProductQuantityGetAPIRequest 分销商查询产品库存 API请求
// taobao.fenxiao.distributor.product.quantity.get
//
// 分销商查询产品库存
type TaobaoFenxiaoDistributorProductQuantityGetAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// SKU ID
	_skuId int64
}

// NewTaobaoFenxiaoDistributorProductQuantityGetRequest 初始化TaobaoFenxiaoDistributorProductQuantityGetAPIRequest对象
func NewTaobaoFenxiaoDistributorProductQuantityGetRequest() *TaobaoFenxiaoDistributorProductQuantityGetAPIRequest {
	return &TaobaoFenxiaoDistributorProductQuantityGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDistributorProductQuantityGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.distributor.product.quantity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDistributorProductQuantityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoDistributorProductQuantityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoDistributorProductQuantityGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoFenxiaoDistributorProductQuantityGetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetSkuId is SkuId Setter
// SKU ID
func (r *TaobaoFenxiaoDistributorProductQuantityGetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoFenxiaoDistributorProductQuantityGetAPIRequest) GetSkuId() int64 {
	return r._skuId
}
