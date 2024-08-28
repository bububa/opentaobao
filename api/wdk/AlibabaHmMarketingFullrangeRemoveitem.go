package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingFullrangeRemoveitem 全场活动删除购品
// alibaba.hm.marketing.fullrange.removeitem
//
// 删除换购商品
func AlibabaHmMarketingFullrangeRemoveitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingFullrangeRemoveitemAPIRequest, resp *wdk.AlibabaHmMarketingFullrangeRemoveitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
