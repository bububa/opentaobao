package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixProjectDistributionQuerybypage 分销项目分页查询项目列表服务
// alibaba.damai.maitix.project.distribution.querybypage
//
// 分销项目分页查询项目列表服务
func AlibabaDamaiMaitixProjectDistributionQuerybypage(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest, resp *maitix.AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
