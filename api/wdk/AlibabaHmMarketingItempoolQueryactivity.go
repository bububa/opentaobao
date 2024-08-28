package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolQueryactivity 查找商品池活动
// alibaba.hm.marketing.itempool.queryactivity
//
// 查找商品池活动
func AlibabaHmMarketingItempoolQueryactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolQueryactivityAPIRequest, resp *wdk.AlibabaHmMarketingItempoolQueryactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
