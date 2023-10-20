package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpWarehouseInventoryQuery AliExpress在仓库存查询API
// aliexpress.ascp.warehouse.inventory.query
//
// AliExpress在仓库存查询查询API
func AliexpressAscpWarehouseInventoryQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpWarehouseInventoryQueryAPIRequest, resp *ascpffo.AliexpressAscpWarehouseInventoryQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
