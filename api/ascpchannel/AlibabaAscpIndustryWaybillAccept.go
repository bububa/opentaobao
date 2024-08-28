package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpIndustryWaybillAccept 商家ERP预推单
// alibaba.ascp.industry.waybill.accept
//
// 商家ERP预推单
func AlibabaAscpIndustryWaybillAccept(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryWaybillAcceptAPIRequest, resp *ascpchannel.AlibabaAscpIndustryWaybillAcceptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
