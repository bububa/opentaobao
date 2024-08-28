package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaAsrDataservicePromotionruleQuery 星巴克优惠规则查询
// alibaba.asr.dataservice.promotionrule.query
//
// 查询优惠规则，例如星巴克查询优惠规则
func AlibabaAsrDataservicePromotionruleQuery(ctx context.Context, clt *core.SDKClient, req *promotion.AlibabaAsrDataservicePromotionruleQueryAPIRequest, resp *promotion.AlibabaAsrDataservicePromotionruleQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
