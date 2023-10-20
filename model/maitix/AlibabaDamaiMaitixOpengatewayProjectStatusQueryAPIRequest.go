package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixopengatewayprojectstatusqueryAPIRequest 分销状态查询接口queryProjectStatusByProjectId API请求
// alibaba.damai.maitix.opengateway.project.status.query
//
// queryProjectStatusByProjectId
type AlibabadamaimaitixopengatewayprojectstatusqueryAPIRequest struct {
	model.Params
	// 入参dto
	_disProjectStatusQueryParam *DisProjectStatusQueryDto
}

// NewAlibabadamaimaitixopengatewayprojectstatusqueryRequest 初始化AlibabadamaimaitixopengatewayprojectstatusqueryAPIRequest对象
func NewAlibabadamaimaitixopengatewayprojectstatusqueryRequest() *AlibabadamaimaitixopengatewayprojectstatusqueryAPIRequest {
	return &AlibabadamaimaitixopengatewayprojectstatusqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixopengatewayprojectstatusqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.opengateway.project.status.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixopengatewayprojectstatusqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixopengatewayprojectstatusqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDisProjectStatusQueryParam is DisProjectStatusQueryParam Setter
// 入参dto
func (r *AlibabadamaimaitixopengatewayprojectstatusqueryAPIRequest) SetDisProjectStatusQueryParam(_disProjectStatusQueryParam *DisProjectStatusQueryDto) error {
	r._disProjectStatusQueryParam = _disProjectStatusQueryParam
	r.Set("dis_project_status_query_param", _disProjectStatusQueryParam)
	return nil
}

// GetDisProjectStatusQueryParam DisProjectStatusQueryParam Getter
func (r AlibabadamaimaitixopengatewayprojectstatusqueryAPIRequest) GetDisProjectStatusQueryParam() *DisProjectStatusQueryDto {
	return r._disProjectStatusQueryParam
}
