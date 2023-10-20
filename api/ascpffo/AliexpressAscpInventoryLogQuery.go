package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpInventoryLogQuery AliExpress库存流水查询API
// aliexpress.ascp.inventory.log.query
//
// AliExpress库存流水查询API
func AliexpressAscpInventoryLogQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpInventoryLogQueryAPIRequest, session string) (*ascpffo.AliexpressAscpInventoryLogQueryAPIResponse, error) {
	var resp ascpffo.AliexpressAscpInventoryLogQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
