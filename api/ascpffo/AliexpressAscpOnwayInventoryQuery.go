package ascpffo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpOnwayInventoryQuery AliExpress在途库存查询API
// aliexpress.ascp.onway.inventory.query
//
// AliExpress在途库存查询API
func AliexpressAscpOnwayInventoryQuery(ctx context.Context, clt *core.SDKClient, req *ascpffo.AliexpressAscpOnwayInventoryQueryAPIRequest, resp *ascpffo.AliexpressAscpOnwayInventoryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
