package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalorderqueryAPIRequest 三方机构查询订单详情接口 API请求
// alibaba.alihealth.medical.order.query
//
// 查询订单详情，包括评价
type AlibabaalihealthmedicalorderqueryAPIRequest struct {
	model.Params
	// 请求入参
	_requestInfo *OrderQueryRequestDto
}

// NewAlibabaalihealthmedicalorderqueryRequest 初始化AlibabaalihealthmedicalorderqueryAPIRequest对象
func NewAlibabaalihealthmedicalorderqueryRequest() *AlibabaalihealthmedicalorderqueryAPIRequest {
	return &AlibabaalihealthmedicalorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestInfo is RequestInfo Setter
// 请求入参
func (r *AlibabaalihealthmedicalorderqueryAPIRequest) SetRequestInfo(_requestInfo *OrderQueryRequestDto) error {
	r._requestInfo = _requestInfo
	r.Set("request_info", _requestInfo)
	return nil
}

// GetRequestInfo RequestInfo Getter
func (r AlibabaalihealthmedicalorderqueryAPIRequest) GetRequestInfo() *OrderQueryRequestDto {
	return r._requestInfo
}
