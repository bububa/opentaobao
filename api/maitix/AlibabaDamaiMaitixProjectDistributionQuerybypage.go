package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixprojectdistributionquerybypage 分销项目分页查询项目列表服务
// alibaba.damai.maitix.project.distribution.querybypage
//
// 分销项目分页查询项目列表服务
func Alibabadamaimaitixprojectdistributionquerybypage(clt *core.SDKClient, req *maitix.AlibabadamaimaitixprojectdistributionquerybypageAPIRequest, session string) (*maitix.AlibabadamaimaitixprojectdistributionquerybypageAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixprojectdistributionquerybypageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
