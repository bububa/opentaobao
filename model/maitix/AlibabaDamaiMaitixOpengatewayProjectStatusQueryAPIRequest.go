package maitix

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest 分销状态查询接口queryProjectStatusByProjectId API请求
// alibaba.damai.maitix.opengateway.project.status.query
//
// queryProjectStatusByProjectId
type AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest struct {
	model.Params
	// 入参dto
	_disProjectStatusQueryParam *DisProjectStatusQueryDto
}

// NewAlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest 初始化AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest对象
func NewAlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest {
	return &AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest) Reset() {
	r._disProjectStatusQueryParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.opengateway.project.status.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDisProjectStatusQueryParam is DisProjectStatusQueryParam Setter
// 入参dto
func (r *AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest) SetDisProjectStatusQueryParam(_disProjectStatusQueryParam *DisProjectStatusQueryDto) error {
	r._disProjectStatusQueryParam = _disProjectStatusQueryParam
	r.Set("dis_project_status_query_param", _disProjectStatusQueryParam)
	return nil
}

// GetDisProjectStatusQueryParam DisProjectStatusQueryParam Getter
func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest) GetDisProjectStatusQueryParam() *DisProjectStatusQueryDto {
	return r._disProjectStatusQueryParam
}

var poolAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest()
	},
}

// GetAlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest 从 sync.Pool 获取 AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest
func GetAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest() *AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest {
	return poolAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest.Get().(*AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest)
}

// ReleaseAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest 将 AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest(v *AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest.Put(v)
}
