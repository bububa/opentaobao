package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingExpirePromotionDelete 短保优惠删除
// alibaba.hm.marketing.expire.promotion.delete
//
// 短保优惠删除
func AlibabaHmMarketingExpirePromotionDelete(clt *core.SDKClient, req *wdk.AlibabaHmMarketingExpirePromotionDeleteAPIRequest, resp *wdk.AlibabaHmMarketingExpirePromotionDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
