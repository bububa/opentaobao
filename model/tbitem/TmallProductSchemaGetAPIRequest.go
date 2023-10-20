package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallproductschemagetAPIRequest 产品信息获取schema获取 API请求
// tmall.product.schema.get
//
// 产品信息获取接口schema形式返回
type TmallproductschemagetAPIRequest struct {
	model.Params
	// 产品编号
	_productId int64
}

// NewTmallproductschemagetRequest 初始化TmallproductschemagetAPIRequest对象
func NewTmallproductschemagetRequest() *TmallproductschemagetAPIRequest {
	return &TmallproductschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallproductschemagetAPIRequest) GetApiMethodName() string {
	return "tmall.product.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallproductschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallproductschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品编号
func (r *TmallproductschemagetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallproductschemagetAPIRequest) GetProductId() int64 {
	return r._productId
}
