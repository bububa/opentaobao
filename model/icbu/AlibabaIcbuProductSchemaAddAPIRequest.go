package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductSchemaAddAPIRequest
（新）商品发布新接口 API请求
alibaba.icbu.product.schema.add

提供发布ICBU商品的入口 */
type AlibabaIcbuProductSchemaAddAPIRequest struct {
	model.Params
	// 发布入参
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaIcbuProductSchemaAddRequest 初始化AlibabaIcbuProductSchemaAddAPIRequest对象
func NewAlibabaIcbuProductSchemaAddRequest() *AlibabaIcbuProductSchemaAddAPIRequest {
	return &AlibabaIcbuProductSchemaAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaAddAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamProductTopPublishRequest Setter
// 发布入参
func (r *AlibabaIcbuProductSchemaAddAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// Get ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaAddAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}
