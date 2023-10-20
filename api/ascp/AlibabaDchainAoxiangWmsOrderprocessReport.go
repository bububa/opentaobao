package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWmsOrderprocessReport 回传发货单流水通知
// alibaba.dchain.aoxiang.wms.orderprocess.report
//
// 回传发货单流水通知
func AlibabaDchainAoxiangWmsOrderprocessReport(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest, resp *ascp.AlibabaDchainAoxiangWmsOrderprocessReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
