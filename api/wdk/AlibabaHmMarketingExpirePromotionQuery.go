package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingExpirePromotionQuery 短保优惠查询
// alibaba.hm.marketing.expire.promotion.query
//
// 短保优惠查询
func AlibabaHmMarketingExpirePromotionQuery(clt *core.SDKClient, req *wdk.AlibabaHmMarketingExpirePromotionQueryAPIRequest, session string) (*wdk.AlibabaHmMarketingExpirePromotionQueryAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingExpirePromotionQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
