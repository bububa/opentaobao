package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionBatchCancel 取消商品分销
// alibaba.dchain.aoxiang.item.distribution.batch.cancel
//
// 取消商品分销
func AlibabaDchainAoxiangItemDistributionBatchCancel(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest, resp *ascp.AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
