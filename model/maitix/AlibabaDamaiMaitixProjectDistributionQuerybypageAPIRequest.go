package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixprojectdistributionquerybypageAPIRequest 分销项目分页查询项目列表服务 API请求
// alibaba.damai.maitix.project.distribution.querybypage
//
// 分销项目分页查询项目列表服务
type AlibabadamaimaitixprojectdistributionquerybypageAPIRequest struct {
	model.Params
	// 入参param
	_param *ProjectPageParam
}

// NewAlibabadamaimaitixprojectdistributionquerybypageRequest 初始化AlibabadamaimaitixprojectdistributionquerybypageAPIRequest对象
func NewAlibabadamaimaitixprojectdistributionquerybypageRequest() *AlibabadamaimaitixprojectdistributionquerybypageAPIRequest {
	return &AlibabadamaimaitixprojectdistributionquerybypageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixprojectdistributionquerybypageAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.project.distribution.querybypage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixprojectdistributionquerybypageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixprojectdistributionquerybypageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参param
func (r *AlibabadamaimaitixprojectdistributionquerybypageAPIRequest) SetParam(_param *ProjectPageParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabadamaimaitixprojectdistributionquerybypageAPIRequest) GetParam() *ProjectPageParam {
	return r._param
}
