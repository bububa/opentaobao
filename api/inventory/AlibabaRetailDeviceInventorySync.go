package inventory

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// AlibabaRetailDeviceInventorySync 库存同步接口
// alibaba.retail.device.inventory.sync
//
// 商库存同步接口
func AlibabaRetailDeviceInventorySync(ctx context.Context, clt *core.SDKClient, req *inventory.AlibabaRetailDeviceInventorySyncAPIRequest, resp *inventory.AlibabaRetailDeviceInventorySyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
