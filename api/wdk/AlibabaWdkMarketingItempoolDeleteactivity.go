package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolDeleteactivity 删除商品池活动
// alibaba.wdk.marketing.itempool.deleteactivity
//
// 删除商品池活动
func AlibabaWdkMarketingItempoolDeleteactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolDeleteactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolDeleteactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
