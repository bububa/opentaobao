package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangIndustryWaybillEdit 服务商编辑运单
// alibaba.dchain.aoxiang.industry.waybill.edit
//
// 服务商编辑运单
func AlibabaDchainAoxiangIndustryWaybillEdit(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangIndustryWaybillEditAPIRequest, session string) (*ascp.AlibabaDchainAoxiangIndustryWaybillEditAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangIndustryWaybillEditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
