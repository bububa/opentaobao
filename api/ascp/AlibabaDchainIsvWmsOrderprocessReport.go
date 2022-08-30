package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainIsvWmsOrderprocessReport 仓作业信息同步
// alibaba.dchain.isv.wms.orderprocess.report
//
// 仓作业信息同步
func AlibabaDchainIsvWmsOrderprocessReport(clt *core.SDKClient, req *ascp.AlibabaDchainIsvWmsOrderprocessReportAPIRequest, session string) (*ascp.AlibabaDchainIsvWmsOrderprocessReportAPIResponse, error) {
	var resp ascp.AlibabaDchainIsvWmsOrderprocessReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
