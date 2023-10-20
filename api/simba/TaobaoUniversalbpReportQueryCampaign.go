package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryCampaign 计划报表查询
// taobao.universalbp.report.query.campaign
//
// 计划报表查询
func TaobaoUniversalbpReportQueryCampaign(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryCampaignAPIRequest, session string) (*simba.TaobaoUniversalbpReportQueryCampaignAPIResponse, error) {
	var resp simba.TaobaoUniversalbpReportQueryCampaignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
