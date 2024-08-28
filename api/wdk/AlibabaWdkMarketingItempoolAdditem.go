package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolAdditem 增加商品池里面的商品
// alibaba.wdk.marketing.itempool.additem
//
// 增加商品池里面的商品
func AlibabaWdkMarketingItempoolAdditem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolAdditemAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolAdditemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
