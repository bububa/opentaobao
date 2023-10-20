package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsync 实仓库存同步
// alibaba.dchain.aoxiang.physics.inventory.batch.upload.async
//
// 实仓库存同步
func AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsync(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest, session string) (*ascp.AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
