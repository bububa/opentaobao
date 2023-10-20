package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixprojectdistributionquerylistAPIRequest 分销项目列表查询（已过时，不推荐使用） API请求
// alibaba.damai.maitix.project.distribution.querylist
//
// 分销项目列表查询接口（已过时，不推荐使用）
type AlibabadamaimaitixprojectdistributionquerylistAPIRequest struct {
	model.Params
}

// NewAlibabadamaimaitixprojectdistributionquerylistRequest 初始化AlibabadamaimaitixprojectdistributionquerylistAPIRequest对象
func NewAlibabadamaimaitixprojectdistributionquerylistRequest() *AlibabadamaimaitixprojectdistributionquerylistAPIRequest {
	return &AlibabadamaimaitixprojectdistributionquerylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixprojectdistributionquerylistAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.project.distribution.querylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixprojectdistributionquerylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixprojectdistributionquerylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}
