package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixprojectdistributionquerylist 分销项目列表查询（已过时，不推荐使用）
// alibaba.damai.maitix.project.distribution.querylist
//
// 分销项目列表查询接口（已过时，不推荐使用）
func Alibabadamaimaitixprojectdistributionquerylist(clt *core.SDKClient, req *maitix.AlibabadamaimaitixprojectdistributionquerylistAPIRequest, session string) (*maitix.AlibabadamaimaitixprojectdistributionquerylistAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixprojectdistributionquerylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
