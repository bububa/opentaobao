package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangIndustryWaybillEdit 服务商编辑运单
// alibaba.dchain.aoxiang.industry.waybill.edit
//
// 服务商编辑运单
func AlibabaDchainAoxiangIndustryWaybillEdit(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangIndustryWaybillEditAPIRequest, resp *ascp.AlibabaDchainAoxiangIndustryWaybillEditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
