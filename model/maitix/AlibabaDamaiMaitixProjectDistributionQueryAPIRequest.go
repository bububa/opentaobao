package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixProjectDistributionQueryAPIRequest 分销单个项目信息查询 API请求
// alibaba.damai.maitix.project.distribution.query
//
// 发布分销项目查询单个项目信息接口
type AlibabaDamaiMaitixProjectDistributionQueryAPIRequest struct {
	model.Params
	// 项目id
	_projectId int64
}

// NewAlibabaDamaiMaitixProjectDistributionQueryRequest 初始化AlibabaDamaiMaitixProjectDistributionQueryAPIRequest对象
func NewAlibabaDamaiMaitixProjectDistributionQueryRequest() *AlibabaDamaiMaitixProjectDistributionQueryAPIRequest {
	return &AlibabaDamaiMaitixProjectDistributionQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixProjectDistributionQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.project.distribution.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixProjectDistributionQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixProjectDistributionQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectId is ProjectId Setter
// 项目id
func (r *AlibabaDamaiMaitixProjectDistributionQueryAPIRequest) SetProjectId(_projectId int64) error {
	r._projectId = _projectId
	r.Set("project_id", _projectId)
	return nil
}

// GetProjectId ProjectId Getter
func (r AlibabaDamaiMaitixProjectDistributionQueryAPIRequest) GetProjectId() int64 {
	return r._projectId
}
