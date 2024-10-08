package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolStairAdditem 商品池阶梯商品添加
// alibaba.wdk.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
func AlibabaWdkMarketingItempoolStairAdditem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolStairAdditemAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolStairAdditemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
