package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkTraceUrlGet 溯源url透出
// alibaba.wdk.trace.url.get
//
// 根据shopId和skuCode返回商品溯源url
func AlibabaWdkTraceUrlGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkTraceUrlGetAPIRequest, resp *wdk.AlibabaWdkTraceUrlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
