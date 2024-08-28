package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingFullrangeDeleteactivity 全场活动删除活动接口
// alibaba.wdk.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
func AlibabaWdkMarketingFullrangeDeleteactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest, resp *wdk.AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
