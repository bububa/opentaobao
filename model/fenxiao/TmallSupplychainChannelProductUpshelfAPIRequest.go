package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallsupplychainchannelproductupshelfAPIRequest 产品上架 API请求
// tmall.supplychain.channel.product.upshelf
//
// 上架渠道产品
type TmallsupplychainchannelproductupshelfAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
}

// NewTmallsupplychainchannelproductupshelfRequest 初始化TmallsupplychainchannelproductupshelfAPIRequest对象
func NewTmallsupplychainchannelproductupshelfRequest() *TmallsupplychainchannelproductupshelfAPIRequest {
	return &TmallsupplychainchannelproductupshelfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallsupplychainchannelproductupshelfAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.upshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallsupplychainchannelproductupshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallsupplychainchannelproductupshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TmallsupplychainchannelproductupshelfAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallsupplychainchannelproductupshelfAPIRequest) GetProductId() int64 {
	return r._productId
}
