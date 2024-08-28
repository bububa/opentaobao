package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolCreateactivity 添加商品池活动
// alibaba.wdk.marketing.itempool.createactivity
//
// 添加商品池活动
func AlibabaWdkMarketingItempoolCreateactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolCreateactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolCreateactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
