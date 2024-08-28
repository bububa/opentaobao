package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolRemoveitem 移除商品池里面的商品
// alibaba.hm.marketing.itempool.removeitem
//
// 移除商品池里面的商品
func AlibabaHmMarketingItempoolRemoveitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolRemoveitemAPIRequest, resp *wdk.AlibabaHmMarketingItempoolRemoveitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
