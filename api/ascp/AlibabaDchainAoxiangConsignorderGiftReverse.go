package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangConsignorderGiftReverse 赠品绑赠回滚
// alibaba.dchain.aoxiang.consignorder.gift.reverse
//
// 赠品绑赠回滚
func AlibabaDchainAoxiangConsignorderGiftReverse(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest, resp *ascp.AlibabaDchainAoxiangConsignorderGiftReverseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
