package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryCampaign 计划报表查询
// taobao.universalbp.report.query.campaign
//
// 计划报表查询
func TaobaoUniversalbpReportQueryCampaign(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryCampaignAPIRequest, resp *simba.TaobaoUniversalbpReportQueryCampaignAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
