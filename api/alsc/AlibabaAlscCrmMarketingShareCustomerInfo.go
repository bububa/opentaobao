package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmMarketingShareCustomerInfo 查询分享营销客户领券信息
// alibaba.alsc.crm.marketing.share.customer.info
//
// 查询分享营销活动的客户领券信息
func AlibabaAlscCrmMarketingShareCustomerInfo(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest, resp *alsc.AlibabaAlscCrmMarketingShareCustomerInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
