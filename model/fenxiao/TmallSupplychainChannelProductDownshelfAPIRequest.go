package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallsupplychainchannelproductdownshelfAPIRequest 产品下架 API请求
// tmall.supplychain.channel.product.downshelf
//
// 产品下架
type TmallsupplychainchannelproductdownshelfAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
}

// NewTmallsupplychainchannelproductdownshelfRequest 初始化TmallsupplychainchannelproductdownshelfAPIRequest对象
func NewTmallsupplychainchannelproductdownshelfRequest() *TmallsupplychainchannelproductdownshelfAPIRequest {
	return &TmallsupplychainchannelproductdownshelfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallsupplychainchannelproductdownshelfAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.downshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallsupplychainchannelproductdownshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallsupplychainchannelproductdownshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TmallsupplychainchannelproductdownshelfAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallsupplychainchannelproductdownshelfAPIRequest) GetProductId() int64 {
	return r._productId
}
