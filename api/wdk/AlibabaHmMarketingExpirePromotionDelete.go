package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingExpirePromotionDelete 短保优惠删除
// alibaba.hm.marketing.expire.promotion.delete
//
// 短保优惠删除
func AlibabaHmMarketingExpirePromotionDelete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingExpirePromotionDeleteAPIRequest, resp *wdk.AlibabaHmMarketingExpirePromotionDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
