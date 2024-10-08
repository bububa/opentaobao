package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangOrderprocessReport 回传仓内作业节点
// alibaba.dchain.aoxiang.orderprocess.report
//
// 回传仓内作业节点
func AlibabaDchainAoxiangOrderprocessReport(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangOrderprocessReportAPIRequest, resp *ascp.AlibabaDchainAoxiangOrderprocessReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
