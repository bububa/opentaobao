package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusChange 物流节点回传
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.change
//
// 物流节点回传
func AlibabaDchainAoxiangIndustryWaybillLogisticstatusChange(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest, resp *ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
