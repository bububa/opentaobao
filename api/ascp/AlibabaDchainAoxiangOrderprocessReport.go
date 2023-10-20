package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangorderprocessreport 回传仓内作业节点
// alibaba.dchain.aoxiang.orderprocess.report
//
// 回传仓内作业节点
func Alibabadchainaoxiangorderprocessreport(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangorderprocessreportAPIRequest, session string) (*ascp.AlibabadchainaoxiangorderprocessreportAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangorderprocessreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
