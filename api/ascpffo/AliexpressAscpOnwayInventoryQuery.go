package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpOnwayInventoryQuery AliExpress在途库存查询API
// aliexpress.ascp.onway.inventory.query
//
// AliExpress在途库存查询API
func AliexpressAscpOnwayInventoryQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpOnwayInventoryQueryAPIRequest, session string) (*ascpffo.AliexpressAscpOnwayInventoryQueryAPIResponse, error) {
	var resp ascpffo.AliexpressAscpOnwayInventoryQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
