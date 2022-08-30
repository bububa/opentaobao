package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangDeliverytemplateQueryAPIRequest 商家运费模板查询 API请求
// alibaba.dchain.aoxiang.deliverytemplate.query
//
// 商家运费模板查询
type AlibabaDchainAoxiangDeliverytemplateQueryAPIRequest struct {
	model.Params
	// 查询运费模板入参
	_queryDeliveryTemplateRequest *QueryDeliveryTemplateRequest
}

// NewAlibabaDchainAoxiangDeliverytemplateQueryRequest 初始化AlibabaDchainAoxiangDeliverytemplateQueryAPIRequest对象
func NewAlibabaDchainAoxiangDeliverytemplateQueryRequest() *AlibabaDchainAoxiangDeliverytemplateQueryAPIRequest {
	return &AlibabaDchainAoxiangDeliverytemplateQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangDeliverytemplateQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.deliverytemplate.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangDeliverytemplateQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQueryDeliveryTemplateRequest is QueryDeliveryTemplateRequest Setter
// 查询运费模板入参
func (r *AlibabaDchainAoxiangDeliverytemplateQueryAPIRequest) SetQueryDeliveryTemplateRequest(_queryDeliveryTemplateRequest *QueryDeliveryTemplateRequest) error {
	r._queryDeliveryTemplateRequest = _queryDeliveryTemplateRequest
	r.Set("query_delivery_template_request", _queryDeliveryTemplateRequest)
	return nil
}

// GetQueryDeliveryTemplateRequest QueryDeliveryTemplateRequest Getter
func (r AlibabaDchainAoxiangDeliverytemplateQueryAPIRequest) GetQueryDeliveryTemplateRequest() *QueryDeliveryTemplateRequest {
	return r._queryDeliveryTemplateRequest
}
