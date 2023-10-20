package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpreportquerycampaign 计划报表查询
// taobao.universalbp.report.query.campaign
//
// 计划报表查询
func Taobaouniversalbpreportquerycampaign(clt *core.SDKClient, req *simba.TaobaouniversalbpreportquerycampaignAPIRequest, session string) (*simba.TaobaouniversalbpreportquerycampaignAPIResponse, error) {
	var resp simba.TaobaouniversalbpreportquerycampaignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
