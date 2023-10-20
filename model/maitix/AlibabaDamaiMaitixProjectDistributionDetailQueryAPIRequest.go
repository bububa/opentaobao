package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixprojectdistributiondetailqueryAPIRequest 大麦分销项目内容详情查询 API请求
// alibaba.damai.maitix.project.distribution.detail.query
//
// 大麦分销项目内容详情查询，提供项目的内容详情
type AlibabadamaimaitixprojectdistributiondetailqueryAPIRequest struct {
	model.Params
	// 项目ID，前提已授权
	_projectId int64
}

// NewAlibabadamaimaitixprojectdistributiondetailqueryRequest 初始化AlibabadamaimaitixprojectdistributiondetailqueryAPIRequest对象
func NewAlibabadamaimaitixprojectdistributiondetailqueryRequest() *AlibabadamaimaitixprojectdistributiondetailqueryAPIRequest {
	return &AlibabadamaimaitixprojectdistributiondetailqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixprojectdistributiondetailqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.project.distribution.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixprojectdistributiondetailqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixprojectdistributiondetailqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectId is ProjectId Setter
// 项目ID，前提已授权
func (r *AlibabadamaimaitixprojectdistributiondetailqueryAPIRequest) SetProjectId(_projectId int64) error {
	r._projectId = _projectId
	r.Set("project_id", _projectId)
	return nil
}

// GetProjectId ProjectId Getter
func (r AlibabadamaimaitixprojectdistributiondetailqueryAPIRequest) GetProjectId() int64 {
	return r._projectId
}
