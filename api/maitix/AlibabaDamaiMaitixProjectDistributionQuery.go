package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixprojectdistributionquery 分销单个项目信息查询
// alibaba.damai.maitix.project.distribution.query
//
// 发布分销项目查询单个项目信息接口
func Alibabadamaimaitixprojectdistributionquery(clt *core.SDKClient, req *maitix.AlibabadamaimaitixprojectdistributionqueryAPIRequest, session string) (*maitix.AlibabadamaimaitixprojectdistributionqueryAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixprojectdistributionqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
