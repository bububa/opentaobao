package ascpffo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpRoItemQuery AliExpress退供单明细查询API
// aliexpress.ascp.ro.item.query
//
// AE仓发 单个退供单明细查询
func AliexpressAscpRoItemQuery(ctx context.Context, clt *core.SDKClient, req *ascpffo.AliexpressAscpRoItemQueryAPIRequest, resp *ascpffo.AliexpressAscpRoItemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
