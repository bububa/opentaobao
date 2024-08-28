package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusChange 物流节点回传
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.change
//
// 物流节点回传
func AlibabaDchainAoxiangIndustryWaybillLogisticstatusChange(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest, resp *ascp.AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
