package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolStairRemoveitem 删除换购活动商品
// alibaba.hm.marketing.itempool.stair.removeitem
//
// 删除换购商品
func AlibabaHmMarketingItempoolStairRemoveitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolStairRemoveitemAPIRequest, resp *wdk.AlibabaHmMarketingItempoolStairRemoveitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
