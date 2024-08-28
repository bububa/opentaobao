package foodscan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/foodscan"
)

// AlibabaFootscanMiniQueryMobilereport 根据scanId查询报告
// alibaba.footscan.mini.query.mobilereport
//
// 根据scanId查询报告
func AlibabaFootscanMiniQueryMobilereport(ctx context.Context, clt *core.SDKClient, req *foodscan.AlibabaFootscanMiniQueryMobilereportAPIRequest, resp *foodscan.AlibabaFootscanMiniQueryMobilereportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
