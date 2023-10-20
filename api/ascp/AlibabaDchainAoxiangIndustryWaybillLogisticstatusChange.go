package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusChange 物流节点回传
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.change
//
// 物流节点回传
func AlibabaDchainAoxiangIndustryWaybillLogisticstatusChange(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest, session string) (*ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
