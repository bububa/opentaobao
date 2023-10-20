package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest 分销项目分页查询项目列表服务 API请求
// alibaba.damai.maitix.project.distribution.querybypage
//
// 分销项目分页查询项目列表服务
type AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest struct {
	model.Params
	// 入参param
	_param *ProjectPageParam
}

// NewAlibabaDamaiMaitixProjectDistributionQuerybypageRequest 初始化AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest对象
func NewAlibabaDamaiMaitixProjectDistributionQuerybypageRequest() *AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest {
	return &AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.project.distribution.querybypage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参param
func (r *AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest) SetParam(_param *ProjectPageParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest) GetParam() *ProjectPageParam {
	return r._param
}
