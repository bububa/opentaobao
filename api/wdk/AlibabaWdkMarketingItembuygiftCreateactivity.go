package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItembuygiftCreateactivity 创建买赠活动
// alibaba.wdk.marketing.itembuygift.createactivity
//
// 创建买赠活动
func AlibabaWdkMarketingItembuygiftCreateactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItembuygiftCreateactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
