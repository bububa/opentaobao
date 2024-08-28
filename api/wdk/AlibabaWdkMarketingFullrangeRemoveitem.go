package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingFullrangeRemoveitem 全场活动删除购品
// alibaba.wdk.marketing.fullrange.removeitem
//
// 删除换购商品
func AlibabaWdkMarketingFullrangeRemoveitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeRemoveitemAPIRequest, resp *wdk.AlibabaWdkMarketingFullrangeRemoveitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
