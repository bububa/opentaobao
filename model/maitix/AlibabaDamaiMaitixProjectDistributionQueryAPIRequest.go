package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixProjectDistributionQueryAPIRequest
分销单个项目信息查询 API请求
alibaba.damai.maitix.project.distribution.query

发布分销项目查询单个项目信息接口 */
type AlibabaDamaiMaitixProjectDistributionQueryAPIRequest struct {
	model.Params
	// 项目id
	_projectId int64
}

// New
