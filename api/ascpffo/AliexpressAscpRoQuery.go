package ascpffo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpRoQuery AliExpress退供单查询API
// aliexpress.ascp.ro.query
//
// AE仓发商家单个退供单查询接口
func AliexpressAscpRoQuery(ctx context.Context, clt *core.SDKClient, req *ascpffo.AliexpressAscpRoQueryAPIRequest, resp *ascpffo.AliexpressAscpRoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
