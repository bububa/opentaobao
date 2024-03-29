package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductUpshelfAPIRequest 产品上架 API请求
// tmall.supplychain.channel.product.upshelf
//
// 上架渠道产品
type TmallSupplychainChannelProductUpshelfAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
}

// NewTmallSupplychainChannelProductUpshelfRequest 初始化TmallSupplychainChannelProductUpshelfAPIRequest对象
func NewTmallSupplychainChannelProductUpshelfRequest() *TmallSupplychainChannelProductUpshelfAPIRequest {
	return &TmallSupplychainChannelProductUpshelfAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallSupplychainChannelProductUpshelfAPIRequest) Reset() {
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductUpshelfAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.upshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductUpshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallSupplychainChannelProductUpshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TmallSupplychainChannelProductUpshelfAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallSupplychainChannelProductUpshelfAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolTmallSupplychainChannelProductUpshelfAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallSupplychainChannelProductUpshelfRequest()
	},
}

// GetTmallSupplychainChannelProductUpshelfRequest 从 sync.Pool 获取 TmallSupplychainChannelProductUpshelfAPIRequest
func GetTmallSupplychainChannelProductUpshelfAPIRequest() *TmallSupplychainChannelProductUpshelfAPIRequest {
	return poolTmallSupplychainChannelProductUpshelfAPIRequest.Get().(*TmallSupplychainChannelProductUpshelfAPIRequest)
}

// ReleaseTmallSupplychainChannelProductUpshelfAPIRequest 将 TmallSupplychainChannelProductUpshelfAPIRequest 放入 sync.Pool
func ReleaseTmallSupplychainChannelProductUpshelfAPIRequest(v *TmallSupplychainChannelProductUpshelfAPIRequest) {
	v.Reset()
	poolTmallSupplychainChannelProductUpshelfAPIRequest.Put(v)
}
