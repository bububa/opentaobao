package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixdistributionexchangepointqueryAPIRequest 分销查询取票点接口 API请求
// alibaba.damai.maitix.distribution.exchangepoint.query
//
// 分销查询取票点接口
type AlibabadamaimaitixdistributionexchangepointqueryAPIRequest struct {
	model.Params
	// 必填-项目id
	_projectId int64
}

// NewAlibabadamaimaitixdistributionexchangepointqueryRequest 初始化AlibabadamaimaitixdistributionexchangepointqueryAPIRequest对象
func NewAlibabadamaimaitixdistributionexchangepointqueryRequest() *AlibabadamaimaitixdistributionexchangepointqueryAPIRequest {
	return &AlibabadamaimaitixdistributionexchangepointqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixdistributionexchangepointqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.distribution.exchangepoint.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixdistributionexchangepointqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixdistributionexchangepointqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectId is ProjectId Setter
// 必填-项目id
func (r *AlibabadamaimaitixdistributionexchangepointqueryAPIRequest) SetProjectId(_projectId int64) error {
	r._projectId = _projectId
	r.Set("project_id", _projectId)
	return nil
}

// GetProjectId ProjectId Getter
func (r AlibabadamaimaitixdistributionexchangepointqueryAPIRequest) GetProjectId() int64 {
	return r._projectId
}
