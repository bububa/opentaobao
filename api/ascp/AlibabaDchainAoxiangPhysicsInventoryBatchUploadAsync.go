package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsync 实仓库存同步
// alibaba.dchain.aoxiang.physics.inventory.batch.upload.async
//
// 实仓库存同步
func AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsync(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest, resp *ascp.AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
