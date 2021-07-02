package maitix

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.distribution.exchangepoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionExchangepointQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
