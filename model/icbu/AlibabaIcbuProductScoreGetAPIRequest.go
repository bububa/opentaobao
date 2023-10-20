package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductscoregetAPIRequest 产品质量分查询 API请求
// alibaba.icbu.product.score.get
//
// 产品质量分查询
type AlibabaicbuproductscoregetAPIRequest struct {
	model.Params
	// 混淆后的商品ID
	_productId string
}

// NewAlibabaicbuproductscoregetRequest 初始化AlibabaicbuproductscoregetAPIRequest对象
func NewAlibabaicbuproductscoregetRequest() *AlibabaicbuproductscoregetAPIRequest {
	return &AlibabaicbuproductscoregetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductscoregetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.score.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductscoregetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductscoregetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 混淆后的商品ID
func (r *AlibabaicbuproductscoregetAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaicbuproductscoregetAPIRequest) GetProductId() string {
	return r._productId
}
