package icbudropshipping

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaOrderFreightCalculate 阿里巴巴下单场景运费方案计算
// alibaba.order.freight.calculate
//
// icbu开展 drop shipping 业务，阿里巴巴下单场景运费方案计算
// alibaba Create order scenario freight calculation
func AlibabaOrderFreightCalculate(ctx context.Context, clt *core.SDKClient, req *icbudropshipping.AlibabaOrderFreightCalculateAPIRequest, resp *icbudropshipping.AlibabaOrderFreightCalculateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
