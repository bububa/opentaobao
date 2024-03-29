package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductDownshelfAPIRequest 产品下架 API请求
// tmall.supplychain.channel.product.downshelf
//
// 产品下架
type TmallSupplychainChannelProductDownshelfAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
}

// NewTmallSupplychainChannelProductDownshelfRequest 初始化TmallSupplychainChannelProductDownshelfAPIRequest对象
func NewTmallSupplychainChannelProductDownshelfRequest() *TmallSupplychainChannelProductDownshelfAPIRequest {
	return &TmallSupplychainChannelProductDownshelfAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallSupplychainChannelProductDownshelfAPIRequest) Reset() {
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductDownshelfAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.downshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductDownshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallSupplychainChannelProductDownshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TmallSupplychainChannelProductDownshelfAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallSupplychainChannelProductDownshelfAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolTmallSupplychainChannelProductDownshelfAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallSupplychainChannelProductDownshelfRequest()
	},
}

// GetTmallSupplychainChannelProductDownshelfRequest 从 sync.Pool 获取 TmallSupplychainChannelProductDownshelfAPIRequest
func GetTmallSupplychainChannelProductDownshelfAPIRequest() *TmallSupplychainChannelProductDownshelfAPIRequest {
	return poolTmallSupplychainChannelProductDownshelfAPIRequest.Get().(*TmallSupplychainChannelProductDownshelfAPIRequest)
}

// ReleaseTmallSupplychainChannelProductDownshelfAPIRequest 将 TmallSupplychainChannelProductDownshelfAPIRequest 放入 sync.Pool
func ReleaseTmallSupplychainChannelProductDownshelfAPIRequest(v *TmallSupplychainChannelProductDownshelfAPIRequest) {
	v.Reset()
	poolTmallSupplychainChannelProductDownshelfAPIRequest.Put(v)
}
