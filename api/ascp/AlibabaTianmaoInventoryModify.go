package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoInventoryModify 阿里巴巴.天猫.aic库存.修改
// alibaba.tianmao.inventory.modify
//
// 阿里巴巴.天猫.aic库存.修改
func AlibabaTianmaoInventoryModify(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaTianmaoInventoryModifyAPIRequest, resp *ascp.AlibabaTianmaoInventoryModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
