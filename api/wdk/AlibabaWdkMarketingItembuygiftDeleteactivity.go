package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItembuygiftDeleteactivity 删除买赠活动
// alibaba.wdk.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
func AlibabaWdkMarketingItembuygiftDeleteactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
