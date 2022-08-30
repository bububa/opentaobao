package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpWarehouseInventoryQuery AliExpress在仓库存查询API
// aliexpress.ascp.warehouse.inventory.query
//
// AliExpress在仓库存查询查询API
func AliexpressAscpWarehouseInventoryQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpWarehouseInventoryQueryAPIRequest, session string) (*ascpffo.AliexpressAscpWarehouseInventoryQueryAPIResponse, error) {
	var resp ascpffo.AliexpressAscpWarehouseInventoryQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
