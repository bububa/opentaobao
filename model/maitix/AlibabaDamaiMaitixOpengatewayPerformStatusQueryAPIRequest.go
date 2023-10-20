package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixopengatewayperformstatusqueryAPIRequest 分销状态查询接口queryPerformStatusByPerformId API请求
// alibaba.damai.maitix.opengateway.perform.status.query
//
// queryPerformStatusByPerformId
type AlibabadamaimaitixopengatewayperformstatusqueryAPIRequest struct {
	model.Params
	// 入参
	_disPerformStatusQueryParam *DisPerformStatusQueryDto
}

// NewAlibabadamaimaitixopengatewayperformstatusqueryRequest 初始化AlibabadamaimaitixopengatewayperformstatusqueryAPIRequest对象
func NewAlibabadamaimaitixopengatewayperformstatusqueryRequest() *AlibabadamaimaitixopengatewayperformstatusqueryAPIRequest {
	return &AlibabadamaimaitixopengatewayperformstatusqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixopengatewayperformstatusqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.opengateway.perform.status.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixopengatewayperformstatusqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixopengatewayperformstatusqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDisPerformStatusQueryParam is DisPerformStatusQueryParam Setter
// 入参
func (r *AlibabadamaimaitixopengatewayperformstatusqueryAPIRequest) SetDisPerformStatusQueryParam(_disPerformStatusQueryParam *DisPerformStatusQueryDto) error {
	r._disPerformStatusQueryParam = _disPerformStatusQueryParam
	r.Set("dis_perform_status_query_param", _disPerformStatusQueryParam)
	return nil
}

// GetDisPerformStatusQueryParam DisPerformStatusQueryParam Getter
func (r AlibabadamaimaitixopengatewayperformstatusqueryAPIRequest) GetDisPerformStatusQueryParam() *DisPerformStatusQueryDto {
	return r._disPerformStatusQueryParam
}
