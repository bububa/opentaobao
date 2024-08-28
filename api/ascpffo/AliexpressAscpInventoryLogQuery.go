package ascpffo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpInventoryLogQuery AliExpress库存流水查询API
// aliexpress.ascp.inventory.log.query
//
// AliExpress库存流水查询API
func AliexpressAscpInventoryLogQuery(ctx context.Context, clt *core.SDKClient, req *ascpffo.AliexpressAscpInventoryLogQueryAPIRequest, resp *ascpffo.AliexpressAscpInventoryLogQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
