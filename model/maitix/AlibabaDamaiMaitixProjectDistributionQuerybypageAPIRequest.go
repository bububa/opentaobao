package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest
分销项目分页查询项目列表服务 API请求
alibaba.damai.maitix.project.distribution.querybypage

分销项目分页查询项目列表服务 */
type AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest struct {
	model.Params
	// 入参param
	_param *ProjectPageParam
}

// New
