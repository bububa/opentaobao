package maitix

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest 分销状态查询接口queryPerformStatusByPerformId API请求
// alibaba.damai.maitix.opengateway.perform.status.query
//
// queryPerformStatusByPerformId
type AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest struct {
	model.Params
	// 入参
	_disPerformStatusQueryParam *DisPerformStatusQueryDto
}

// NewAlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest 初始化AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest对象
func NewAlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest {
	return &AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest) Reset() {
	r._disPerformStatusQueryParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.opengateway.perform.status.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDisPerformStatusQueryParam is DisPerformStatusQueryParam Setter
// 入参
func (r *AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest) SetDisPerformStatusQueryParam(_disPerformStatusQueryParam *DisPerformStatusQueryDto) error {
	r._disPerformStatusQueryParam = _disPerformStatusQueryParam
	r.Set("dis_perform_status_query_param", _disPerformStatusQueryParam)
	return nil
}

// GetDisPerformStatusQueryParam DisPerformStatusQueryParam Getter
func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest) GetDisPerformStatusQueryParam() *DisPerformStatusQueryDto {
	return r._disPerformStatusQueryParam
}

var poolAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest()
	},
}

// GetAlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest 从 sync.Pool 获取 AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest
func GetAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest() *AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest {
	return poolAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest.Get().(*AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest)
}

// ReleaseAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest 将 AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest(v *AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest.Put(v)
}
