package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixProjectDistributionQuery 分销单个项目信息查询
// alibaba.damai.maitix.project.distribution.query
//
// 发布分销项目查询单个项目信息接口
func AlibabaDamaiMaitixProjectDistributionQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixProjectDistributionQueryAPIRequest, session string) (*maitix.AlibabaDamaiMaitixProjectDistributionQueryAPIResponse, error) {
	var resp maitix.AlibabaDamaiMaitixProjectDistributionQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
