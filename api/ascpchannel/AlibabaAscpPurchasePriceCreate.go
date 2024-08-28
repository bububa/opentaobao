package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpPurchasePriceCreate ascp采购价写入接口
// alibaba.ascp.purchase.price.create
//
// 供应链平台采购价创建或修改接口
func AlibabaAscpPurchasePriceCreate(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpPurchasePriceCreateAPIRequest, resp *ascpchannel.AlibabaAscpPurchasePriceCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
