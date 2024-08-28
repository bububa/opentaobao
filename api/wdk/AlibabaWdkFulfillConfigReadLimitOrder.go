package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillConfigReadLimitOrder 根据仓code查询仓限单配置
// alibaba.wdk.fulfill.config.read.limit.order
//
// 根据仓code查询仓限单配置
func AlibabaWdkFulfillConfigReadLimitOrder(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkFulfillConfigReadLimitOrderAPIRequest, resp *wdk.AlibabaWdkFulfillConfigReadLimitOrderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
