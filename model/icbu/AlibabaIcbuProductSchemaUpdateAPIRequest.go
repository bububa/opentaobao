package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductSchemaUpdateAPIRequest （新）商品发布增量更新接口 API请求
// alibaba.icbu.product.schema.update
//
// 商品更新接口，方式为增量更新，只更新传入的字段
type AlibabaIcbuProductSchemaUpdateAPIRequest struct {
	model.Params
	// 发布入参
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaIcbuProductSchemaUpdateRequest 初始化AlibabaIcbuProductSchemaUpdateAPIRequest对象
func NewAlibabaIcbuProductSchemaUpdateRequest() *AlibabaIcbuProductSchemaUpdateAPIRequest {
	return &AlibabaIcbuProductSchemaUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamProductTopPublishRequest Setter
// 发布入参
func (r *AlibabaIcbuProductSchemaUpdateAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// Get ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaUpdateAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}
