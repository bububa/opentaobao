package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingFullrangeAddexchangeitem 全场增加换购品
// alibaba.hm.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
func AlibabaHmMarketingFullrangeAddexchangeitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest, resp *wdk.AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
