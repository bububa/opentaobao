package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingExpirePromotionQuery 短保优惠查询
// alibaba.wdk.marketing.expire.promotion.query
//
// 短保优惠查询
func AlibabaWdkMarketingExpirePromotionQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingExpirePromotionQueryAPIRequest, resp *wdk.AlibabaWdkMarketingExpirePromotionQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
