package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallproductupdateschemagetAPIRequest 产品更新规则获取接口 API请求
// tmall.product.update.schema.get
//
// 获取用户更新产品的规则
type TmallproductupdateschemagetAPIRequest struct {
	model.Params
	// 产品编号
	_productId int64
}

// NewTmallproductupdateschemagetRequest 初始化TmallproductupdateschemagetAPIRequest对象
func NewTmallproductupdateschemagetRequest() *TmallproductupdateschemagetAPIRequest {
	return &TmallproductupdateschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallproductupdateschemagetAPIRequest) GetApiMethodName() string {
	return "tmall.product.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallproductupdateschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallproductupdateschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品编号
func (r *TmallproductupdateschemagetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallproductupdateschemagetAPIRequest) GetProductId() int64 {
	return r._productId
}
