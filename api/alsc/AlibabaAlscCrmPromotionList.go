package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmPromotionList 获取促销规则列表
// alibaba.alsc.crm.promotion.list
//
// 获取品牌的促销规则列表
func AlibabaAlscCrmPromotionList(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmPromotionListAPIRequest, resp *alsc.AlibabaAlscCrmPromotionListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
