package lstmarketing

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstmarketing"
)

// AlibabaLstMarketingQuerybyorderid 根据订单查询营销信息
// alibaba.lst.marketing.querybyorderid
//
// 根据订单查询营销信息
func AlibabaLstMarketingQuerybyorderid(ctx context.Context, clt *core.SDKClient, req *lstmarketing.AlibabaLstMarketingQuerybyorderidAPIRequest, resp *lstmarketing.AlibabaLstMarketingQuerybyorderidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
