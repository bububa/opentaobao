package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpIndustryUopSupplierConsignoder 商家推单
// alibaba.ascp.industry.uop.supplier.consignoder
//
// 商家推单
func AlibabaAscpIndustryUopSupplierConsignoder(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryUopSupplierConsignoderAPIRequest, resp *ascpchannel.AlibabaAscpIndustryUopSupplierConsignoderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
