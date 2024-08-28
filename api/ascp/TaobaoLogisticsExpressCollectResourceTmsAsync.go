package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressCollectResourceTmsAsync 配服务商揽收能力同步接口
// taobao.logistics.express.collect.resource.tms.async
//
// 配服务商揽收能力同步接口
func TaobaoLogisticsExpressCollectResourceTmsAsync(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest, resp *ascp.TaobaoLogisticsExpressCollectResourceTmsAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
