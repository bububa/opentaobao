package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWarehouseStatusUpdate 启用/停用仓资源
// alibaba.dchain.aoxiang.warehouse.status.update
//
// 启用/停用仓资源
func AlibabaDchainAoxiangWarehouseStatusUpdate(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWarehouseStatusUpdateAPIRequest, resp *ascp.AlibabaDchainAoxiangWarehouseStatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
