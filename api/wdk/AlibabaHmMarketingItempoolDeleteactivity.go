package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolDeleteactivity 删除商品池活动
// alibaba.hm.marketing.itempool.deleteactivity
//
// 删除商品池活动
func AlibabaHmMarketingItempoolDeleteactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolDeleteactivityAPIRequest, resp *wdk.AlibabaHmMarketingItempoolDeleteactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
