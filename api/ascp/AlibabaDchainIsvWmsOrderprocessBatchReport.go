package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainIsvWmsOrderprocessBatchReport 仓作业信息批量同步
// alibaba.dchain.isv.wms.orderprocess.batch.report
//
// 仓作业信息批量同步
func AlibabaDchainIsvWmsOrderprocessBatchReport(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainIsvWmsOrderprocessBatchReportAPIRequest, resp *ascp.AlibabaDchainIsvWmsOrderprocessBatchReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
