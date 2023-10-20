package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangIndustryWaybillEdit 服务商编辑运单
// alibaba.dchain.aoxiang.industry.waybill.edit
//
// 服务商编辑运单
func AlibabaDchainAoxiangIndustryWaybillEdit(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangIndustryWaybillEditAPIRequest, resp *ascp.AlibabaDchainAoxiangIndustryWaybillEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
