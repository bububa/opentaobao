package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest 大麦分销项目内容详情查询 API请求
// alibaba.damai.maitix.project.distribution.detail.query
//
// 大麦分销项目内容详情查询，提供项目的内容详情
type AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest struct {
	model.Params
	// 项目ID，前提已授权
	_projectId int64
}

// NewAlibabaDamaiMaitixProjectDistributionDetailQueryRequest 初始化AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest对象
func NewAlibabaDamaiMaitixProjectDistributionDetailQueryRequest() *AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest {
	return &AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.project.distribution.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProjectId is ProjectId Setter
// 项目ID，前提已授权
func (r *AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest) SetProjectId(_projectId int64) error {
	r._projectId = _projectId
	r.Set("project_id", _projectId)
	return nil
}

// GetProjectId ProjectId Getter
func (r AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest) GetProjectId() int64 {
	return r._projectId
}
