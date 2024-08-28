package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingFullrangeQueryactivity 全场活动查询活动
// alibaba.hm.marketing.fullrange.queryactivity
//
// 全场活动查询活动
func AlibabaHmMarketingFullrangeQueryactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingFullrangeQueryactivityAPIRequest, resp *wdk.AlibabaHmMarketingFullrangeQueryactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
