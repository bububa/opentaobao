package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscpurchaseservicedefinitionparamqueryAPIRequest 查询采购服务定义参数信息 API请求
// alibaba.ssc.purchase.servicedefinition.param.query
//
// 查询采购服务定义参数信息
type AlibabasscpurchaseservicedefinitionparamqueryAPIRequest struct {
	model.Params
	// 服务产品id
	_productId int64
}

// NewAlibabasscpurchaseservicedefinitionparamqueryRequest 初始化AlibabasscpurchaseservicedefinitionparamqueryAPIRequest对象
func NewAlibabasscpurchaseservicedefinitionparamqueryRequest() *AlibabasscpurchaseservicedefinitionparamqueryAPIRequest {
	return &AlibabasscpurchaseservicedefinitionparamqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscpurchaseservicedefinitionparamqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.purchase.servicedefinition.param.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscpurchaseservicedefinitionparamqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscpurchaseservicedefinitionparamqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 服务产品id
func (r *AlibabasscpurchaseservicedefinitionparamqueryAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabasscpurchaseservicedefinitionparamqueryAPIRequest) GetProductId() int64 {
	return r._productId
}
