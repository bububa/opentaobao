package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingFullrangeQueryactivity 全场活动查询活动
// alibaba.wdk.marketing.fullrange.queryactivity
//
// 全场活动查询活动
func AlibabaWdkMarketingFullrangeQueryactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeQueryactivityAPIRequest, resp *wdk.AlibabaWdkMarketingFullrangeQueryactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
