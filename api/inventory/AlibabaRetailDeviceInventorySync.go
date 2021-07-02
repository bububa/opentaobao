package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// AlibabaRetailDeviceInventorySync 库存同步接口
// alibaba.retail.device.inventory.sync
//
// 商库存同步接口
func AlibabaRetailDeviceInventorySync(clt *core.SDKClient, req *inventory.AlibabaRetailDeviceInventorySyncAPIRequest, session string) (*inventory.AlibabaRetailDeviceInventorySyncAPIResponse, error) {
	var resp inventory.AlibabaRetailDeviceInventorySyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
