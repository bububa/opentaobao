package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixprojectdistributionqueryAPIRequest 分销单个项目信息查询 API请求
// alibaba.damai.maitix.project.distribution.query
//
// 发布分销项目查询单个项目信息接口
type AlibabadamaimaitixprojectdistributionqueryAPIRequest struct {
	model.Params
	// 项目id
	_projectId int64
}

// NewAlibabadamaimaitixprojectdistributionqueryRequest 初始化AlibabadamaimaitixprojectdistributionqueryAPIRequest对象
func NewAlibabadamaimaitixprojectdistributionqueryRequest() *AlibabadamaimaitixprojectdistributionqueryAPIRequest {
	return &AlibabadamaimaitixprojectdistributionqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixprojectdistributionqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.project.distribution.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixprojectdistributionqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixprojectdistributionqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectId is ProjectId Setter
// 项目id
func (r *AlibabadamaimaitixprojectdistributionqueryAPIRequest) SetProjectId(_projectId int64) error {
	r._projectId = _projectId
	r.Set("project_id", _projectId)
	return nil
}

// GetProjectId ProjectId Getter
func (r AlibabadamaimaitixprojectdistributionqueryAPIRequest) GetProjectId() int64 {
	return r._projectId
}
