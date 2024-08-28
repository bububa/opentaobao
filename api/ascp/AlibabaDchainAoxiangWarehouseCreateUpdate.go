package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWarehouseCreateUpdate 仓库信息同步
// alibaba.dchain.aoxiang.warehouse.create.update
//
// 仓库信息同步
func AlibabaDchainAoxiangWarehouseCreateUpdate(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWarehouseCreateUpdateAPIRequest, resp *ascp.AlibabaDchainAoxiangWarehouseCreateUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
