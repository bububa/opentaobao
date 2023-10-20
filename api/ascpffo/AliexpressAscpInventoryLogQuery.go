package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpInventoryLogQuery AliExpress库存流水查询API
// aliexpress.ascp.inventory.log.query
//
// AliExpress库存流水查询API
func AliexpressAscpInventoryLogQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpInventoryLogQueryAPIRequest, resp *ascpffo.AliexpressAscpInventoryLogQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
