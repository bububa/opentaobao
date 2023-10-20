package maitix

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest 分销查询取票点接口 API请求
// alibaba.damai.maitix.distribution.exchangepoint.query
//
// 分销查询取票点接口
type AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest struct {
	model.Params
	// 必填-项目id
	_projectId int64
}

// NewAlibabaDamaiMaitixDistributionExchangepointQueryRequest 初始化AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest对象
func NewAlibabaDamaiMaitixDistributionExchangepointQueryRequest() *AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest {
	return &AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest) Reset() {
	r._projectId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.distribution.exchangepoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectId is ProjectId Setter
// 必填-项目id
func (r *AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest) SetProjectId(_projectId int64) error {
	r._projectId = _projectId
	r.Set("project_id", _projectId)
	return nil
}

// GetProjectId ProjectId Getter
func (r AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest) GetProjectId() int64 {
	return r._projectId
}

var poolAlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixDistributionExchangepointQueryRequest()
	},
}

// GetAlibabaDamaiMaitixDistributionExchangepointQueryRequest 从 sync.Pool 获取 AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest
func GetAlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest() *AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest {
	return poolAlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest.Get().(*AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest)
}

// ReleaseAlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest 将 AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest(v *AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest.Put(v)
}
