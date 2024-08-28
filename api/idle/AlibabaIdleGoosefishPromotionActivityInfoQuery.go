package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleGoosefishPromotionActivityInfoQuery 闲鱼三方活动参与信息查询
// alibaba.idle.goosefish.promotion.activity.info.query
//
// 闲鱼三方活动参与信息查询
func AlibabaIdleGoosefishPromotionActivityInfoQuery(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest, resp *idle.AlibabaIdleGoosefishPromotionActivityInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
