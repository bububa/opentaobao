package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemBatchDeleteAsync 货品与组合货品删除
// alibaba.dchain.aoxiang.item.batch.delete.async
//
// 货品与组合货品删除
func AlibabaDchainAoxiangItemBatchDeleteAsync(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemBatchDeleteAsyncAPIRequest, resp *ascp.AlibabaDchainAoxiangItemBatchDeleteAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
