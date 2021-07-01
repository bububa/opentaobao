package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkMarketingExpirePromotionQuery
短保优惠查询
alibaba.wdk.marketing.expire.promotion.query

短保优惠查询 */
func AlibabaWdkMarketingExpirePromotionQuery(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingExpirePromotionQueryAPIRequest, session string) (*wdk.AlibabaWdkMarketingExpirePromotionQueryAPIResponse, error) {
	var resp wdk.AlibabaWdkMarketingExpirePromotionQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
