package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaIfpFulfillWarehouseTokenQuery 同城令牌打印接口
// alibaba.ifp.fulfill.warehouse.token.query
//
// 用于仓内作业打印包裹信息
func AlibabaIfpFulfillWarehouseTokenQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaIfpFulfillWarehouseTokenQueryAPIRequest, resp *wdk.AlibabaIfpFulfillWarehouseTokenQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
