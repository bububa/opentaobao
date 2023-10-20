package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingExpirePromotionQuery 短保优惠查询
// alibaba.hm.marketing.expire.promotion.query
//
// 短保优惠查询
func AlibabaHmMarketingExpirePromotionQuery(clt *core.SDKClient, req *wdk.AlibabaHmMarketingExpirePromotionQueryAPIRequest, resp *wdk.AlibabaHmMarketingExpirePromotionQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
