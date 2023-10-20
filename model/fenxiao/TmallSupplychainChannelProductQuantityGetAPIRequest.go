package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductQuantityGetAPIRequest 渠道库存查询接口 API请求
// tmall.supplychain.channel.product.quantity.get
//
// 渠道库存查询接口
type TmallSupplychainChannelProductQuantityGetAPIRequest struct {
	model.Params
	// 产品数字ID
	_productId int64
	// SKU ID
	_skuId int64
}

// NewTmallSupplychainChannelProductQuantityGetRequest 初始化TmallSupplychainChannelProductQuantityGetAPIRequest对象
func NewTmallSupplychainChannelProductQuantityGetRequest() *TmallSupplychainChannelProductQuantityGetAPIRequest {
	return &TmallSupplychainChannelProductQuantityGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallSupplychainChannelProductQuantityGetAPIRequest) Reset() {
	r._productId = 0
	r._skuId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductQuantityGetAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.quantity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductQuantityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallSupplychainChannelProductQuantityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品数字ID
func (r *TmallSupplychainChannelProductQuantityGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallSupplychainChannelProductQuantityGetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetSkuId is SkuId Setter
// SKU ID
func (r *TmallSupplychainChannelProductQuantityGetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TmallSupplychainChannelProductQuantityGetAPIRequest) GetSkuId() int64 {
	return r._skuId
}

var poolTmallSupplychainChannelProductQuantityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallSupplychainChannelProductQuantityGetRequest()
	},
}

// GetTmallSupplychainChannelProductQuantityGetRequest 从 sync.Pool 获取 TmallSupplychainChannelProductQuantityGetAPIRequest
func GetTmallSupplychainChannelProductQuantityGetAPIRequest() *TmallSupplychainChannelProductQuantityGetAPIRequest {
	return poolTmallSupplychainChannelProductQuantityGetAPIRequest.Get().(*TmallSupplychainChannelProductQuantityGetAPIRequest)
}

// ReleaseTmallSupplychainChannelProductQuantityGetAPIRequest 将 TmallSupplychainChannelProductQuantityGetAPIRequest 放入 sync.Pool
func ReleaseTmallSupplychainChannelProductQuantityGetAPIRequest(v *TmallSupplychainChannelProductQuantityGetAPIRequest) {
	v.Reset()
	poolTmallSupplychainChannelProductQuantityGetAPIRequest.Put(v)
}
