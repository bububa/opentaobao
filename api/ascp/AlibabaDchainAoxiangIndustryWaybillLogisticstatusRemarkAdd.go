package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAdd 客服增加备注
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.remark.add
//
// 客服增加备注
func AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAdd(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest, session string) (*ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
