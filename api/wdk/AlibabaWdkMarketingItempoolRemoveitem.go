package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolRemoveitem 移除商品池里面的商品
// alibaba.wdk.marketing.itempool.removeitem
//
// 移除商品池里面的商品
func AlibabaWdkMarketingItempoolRemoveitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolRemoveitemAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolRemoveitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
