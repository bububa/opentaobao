package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangIndustryWaybillCreate 服务商开运单
// alibaba.dchain.aoxiang.industry.waybill.create
//
// 服务商开运单
func AlibabaDchainAoxiangIndustryWaybillCreate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
