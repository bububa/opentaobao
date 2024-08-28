package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolQueryitems 查询商品池活动下的商品
// alibaba.wdk.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
func AlibabaWdkMarketingItempoolQueryitems(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolQueryitemsAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolQueryitemsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
