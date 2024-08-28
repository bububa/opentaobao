package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemBatchUpdateAsync 货品新建/更新接口
// alibaba.dchain.aoxiang.item.batch.update.async
//
// 货品新建/更新接口
func AlibabaDchainAoxiangItemBatchUpdateAsync(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemBatchUpdateAsyncAPIRequest, resp *ascp.AlibabaDchainAoxiangItemBatchUpdateAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
