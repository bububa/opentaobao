package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpIndustryWaybillPreAccept 商家ERP预推单
// alibaba.ascp.industry.waybill.pre.accept
//
// 商家ERP预推单
func AlibabaAscpIndustryWaybillPreAccept(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryWaybillPreAcceptAPIRequest, resp *ascpchannel.AlibabaAscpIndustryWaybillPreAcceptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
