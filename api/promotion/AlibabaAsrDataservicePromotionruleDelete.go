package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaAsrDataservicePromotionruleDelete 优惠规则删除
// alibaba.asr.dataservice.promotionrule.delete
//
// 删除优惠规则，例如星巴克删除优惠规则
func AlibabaAsrDataservicePromotionruleDelete(ctx context.Context, clt *core.SDKClient, req *promotion.AlibabaAsrDataservicePromotionruleDeleteAPIRequest, resp *promotion.AlibabaAsrDataservicePromotionruleDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
