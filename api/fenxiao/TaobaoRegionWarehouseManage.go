package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoRegionWarehouseManage 编辑仓库覆盖范围
// taobao.region.warehouse.manage
//
// 编辑仓库覆盖范围
func TaobaoRegionWarehouseManage(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoRegionWarehouseManageAPIRequest, resp *fenxiao.TaobaoRegionWarehouseManageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
