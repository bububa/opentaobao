package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangChannelInventoryBatchUpload ERP全量同步销售库存数量
// alibaba.dchain.aoxiang.channel.inventory.batch.upload
//
// ERP全量同步销售库存数量
func AlibabaDchainAoxiangChannelInventoryBatchUpload(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest, resp *ascp.AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
