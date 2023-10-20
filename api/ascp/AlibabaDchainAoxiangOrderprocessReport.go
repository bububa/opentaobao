package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangOrderprocessReport 回传仓内作业节点
// alibaba.dchain.aoxiang.orderprocess.report
//
// 回传仓内作业节点
func AlibabaDchainAoxiangOrderprocessReport(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangOrderprocessReportAPIRequest, session string) (*ascp.AlibabaDchainAoxiangOrderprocessReportAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangOrderprocessReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
