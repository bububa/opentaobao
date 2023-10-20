package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangIndustryWaybillCreate 服务商开运单
// alibaba.dchain.aoxiang.industry.waybill.create
//
// 服务商开运单
func AlibabaDchainAoxiangIndustryWaybillCreate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest, resp *ascp.AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
