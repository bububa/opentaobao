package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItembuygiftQueryactivity 查询买赠活动
// alibaba.hm.marketing.itembuygift.queryactivity
//
// 查询买赠活动
func AlibabaHmMarketingItembuygiftQueryactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItembuygiftQueryactivityAPIRequest, resp *wdk.AlibabaHmMarketingItembuygiftQueryactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
