package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryCrowd 人群报表查询
// taobao.universalbp.report.query.crowd
//
// 人群报表查询
func TaobaoUniversalbpReportQueryCrowd(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryCrowdAPIRequest, resp *simba.TaobaoUniversalbpReportQueryCrowdAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
