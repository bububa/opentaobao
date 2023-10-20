package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAdd 客服增加备注
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.remark.add
//
// 客服增加备注
func AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAdd(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest, resp *ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
