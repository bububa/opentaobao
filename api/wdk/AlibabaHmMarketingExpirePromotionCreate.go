package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingExpirePromotionCreate 短保优惠创建
// alibaba.hm.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
func AlibabaHmMarketingExpirePromotionCreate(clt *core.SDKClient, req *wdk.AlibabaHmMarketingExpirePromotionCreateAPIRequest, session string) (*wdk.AlibabaHmMarketingExpirePromotionCreateAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingExpirePromotionCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
