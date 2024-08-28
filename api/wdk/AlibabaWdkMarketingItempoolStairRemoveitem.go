package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolStairRemoveitem 删除换购活动商品
// alibaba.wdk.marketing.itempool.stair.removeitem
//
// 删除换购商品
func AlibabaWdkMarketingItempoolStairRemoveitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolStairRemoveitemAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
