package ascpffo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpItemQuery AliExpress货品查询查询API
// aliexpress.ascp.item.query
//
// AE货品查询API
func AliexpressAscpItemQuery(ctx context.Context, clt *core.SDKClient, req *ascpffo.AliexpressAscpItemQueryAPIRequest, resp *ascpffo.AliexpressAscpItemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
