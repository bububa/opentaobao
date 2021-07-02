package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalOrderQueryAPIRequest 三方机构查询订单详情接口 API请求
// alibaba.alihealth.medical.order.query
//
// 查询订单详情，包括评价
type AlibabaAlihealthMedicalOrderQueryAPIRequest struct {
	model.Params
	// 请求入参
	_requestInfo *OrderQueryRequestDto
}

// NewAlibabaAlihealthMedicalOrderQueryRequest 初始化AlibabaAlihealthMedicalOrderQueryAPIRequest对象
func NewAlibabaAlihealthMedicalOrderQueryRequest() *AlibabaAlihealthMedicalOrderQueryAPIRequest {
	return &AlibabaAlihealthMedicalOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RequestInfo Setter
// 请求入参
func (r *AlibabaAlihealthMedicalOrderQueryAPIRequest) SetRequestInfo(_requestInfo *OrderQueryRequestDto) error {
	r._requestInfo = _requestInfo
	r.Set("request_info", _requestInfo)
	return nil
}

// Get RequestInfo Getter
func (r AlibabaAlihealthMedicalOrderQueryAPIRequest) GetRequestInfo() *OrderQueryRequestDto {
	return r._requestInfo
}
