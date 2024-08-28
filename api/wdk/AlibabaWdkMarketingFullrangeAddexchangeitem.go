package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingFullrangeAddexchangeitem 全场增加换购品
// alibaba.wdk.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
func AlibabaWdkMarketingFullrangeAddexchangeitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest, resp *wdk.AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
