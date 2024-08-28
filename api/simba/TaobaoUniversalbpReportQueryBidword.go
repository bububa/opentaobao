package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryBidword 关键词报表查询
// taobao.universalbp.report.query.bidword
//
// 关键词报表查询
func TaobaoUniversalbpReportQueryBidword(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryBidwordAPIRequest, resp *simba.TaobaoUniversalbpReportQueryBidwordAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
