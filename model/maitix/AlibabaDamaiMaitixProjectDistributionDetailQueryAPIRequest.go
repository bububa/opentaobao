package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest
大麦分销项目内容详情查询 API请求
alibaba.damai.maitix.project.distribution.detail.query

大麦分销项目内容详情查询，提供项目的内容详情 */
type AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest struct {
	model.Params
	// 项目ID，前提已授权
	_projectId int64
}

// New
