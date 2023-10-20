package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixProjectDistributionQuery 分销单个项目信息查询
// alibaba.damai.maitix.project.distribution.query
//
// 发布分销项目查询单个项目信息接口
func AlibabaDamaiMaitixProjectDistributionQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixProjectDistributionQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixProjectDistributionQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
