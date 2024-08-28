package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionUpdate 更新商品分销内容
// alibaba.dchain.aoxiang.item.distribution.update
//
// 更新商品分销内容
func AlibabaDchainAoxiangItemDistributionUpdate(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionUpdateAPIRequest, resp *ascp.AlibabaDchainAoxiangItemDistributionUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
