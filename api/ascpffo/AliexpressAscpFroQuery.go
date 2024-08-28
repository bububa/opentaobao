package ascpffo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpFroQuery AliExpress销退单查询API
// aliexpress.ascp.fro.query
//
// AE履约销退单查询接口
func AliexpressAscpFroQuery(ctx context.Context, clt *core.SDKClient, req *ascpffo.AliexpressAscpFroQueryAPIRequest, resp *ascpffo.AliexpressAscpFroQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
