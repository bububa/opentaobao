package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoInventoryQuery 阿里巴巴.天猫.aic库存.查询
// alibaba.tianmao.inventory.query
//
// 阿里巴巴.天猫.aic库存.查询
func AlibabaTianmaoInventoryQuery(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaTianmaoInventoryQueryAPIRequest, resp *ascp.AlibabaTianmaoInventoryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
