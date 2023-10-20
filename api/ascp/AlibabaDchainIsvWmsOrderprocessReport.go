package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainIsvWmsOrderprocessReport 仓作业信息同步
// alibaba.dchain.isv.wms.orderprocess.report
//
// 仓作业信息同步
func AlibabaDchainIsvWmsOrderprocessReport(clt *core.SDKClient, req *ascp.AlibabaDchainIsvWmsOrderprocessReportAPIRequest, resp *ascp.AlibabaDchainIsvWmsOrderprocessReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
