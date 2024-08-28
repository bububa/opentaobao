package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryCampaign 计划报表查询
// taobao.universalbp.report.query.campaign
//
// 计划报表查询
func TaobaoUniversalbpReportQueryCampaign(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryCampaignAPIRequest, resp *simba.TaobaoUniversalbpReportQueryCampaignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
