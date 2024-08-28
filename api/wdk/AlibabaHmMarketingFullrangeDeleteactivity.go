package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingFullrangeDeleteactivity 全场活动删除活动接口
// alibaba.hm.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
func AlibabaHmMarketingFullrangeDeleteactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingFullrangeDeleteactivityAPIRequest, resp *wdk.AlibabaHmMarketingFullrangeDeleteactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
