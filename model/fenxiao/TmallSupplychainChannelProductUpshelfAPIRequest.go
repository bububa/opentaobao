package fenxiao

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductUpshelfAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.upshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductUpshelfAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
