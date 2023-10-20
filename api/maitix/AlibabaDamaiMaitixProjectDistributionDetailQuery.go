package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixprojectdistributiondetailquery 大麦分销项目内容详情查询
// alibaba.damai.maitix.project.distribution.detail.query
//
// 大麦分销项目内容详情查询，提供项目的内容详情
func Alibabadamaimaitixprojectdistributiondetailquery(clt *core.SDKClient, req *maitix.AlibabadamaimaitixprojectdistributiondetailqueryAPIRequest, session string) (*maitix.AlibabadamaimaitixprojectdistributiondetailqueryAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixprojectdistributiondetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
