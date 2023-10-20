package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangdeliverytemplatequeryAPIRequest 商家运费模板查询 API请求
// alibaba.dchain.aoxiang.deliverytemplate.query
//
// 商家运费模板查询
type AlibabadchainaoxiangdeliverytemplatequeryAPIRequest struct {
	model.Params
	// 查询运费模板入参
	_queryDeliveryTemplateRequest *QueryDeliveryTemplateRequest
}

// NewAlibabadchainaoxiangdeliverytemplatequeryRequest 初始化AlibabadchainaoxiangdeliverytemplatequeryAPIRequest对象
func NewAlibabadchainaoxiangdeliverytemplatequeryRequest() *AlibabadchainaoxiangdeliverytemplatequeryAPIRequest {
	return &AlibabadchainaoxiangdeliverytemplatequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangdeliverytemplatequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.deliverytemplate.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangdeliverytemplatequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangdeliverytemplatequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDeliveryTemplateRequest is QueryDeliveryTemplateRequest Setter
// 查询运费模板入参
func (r *AlibabadchainaoxiangdeliverytemplatequeryAPIRequest) SetQueryDeliveryTemplateRequest(_queryDeliveryTemplateRequest *QueryDeliveryTemplateRequest) error {
	r._queryDeliveryTemplateRequest = _queryDeliveryTemplateRequest
	r.Set("query_delivery_template_request", _queryDeliveryTemplateRequest)
	return nil
}

// GetQueryDeliveryTemplateRequest QueryDeliveryTemplateRequest Getter
func (r AlibabadchainaoxiangdeliverytemplatequeryAPIRequest) GetQueryDeliveryTemplateRequest() *QueryDeliveryTemplateRequest {
	return r._queryDeliveryTemplateRequest
}
