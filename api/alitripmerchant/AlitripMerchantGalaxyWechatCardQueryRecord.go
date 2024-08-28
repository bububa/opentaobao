package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatCardQueryRecord 微信会员卡领取记录查询
// alitrip.merchant.galaxy.wechat.card.query.record
//
// 微信会员卡领取记录查询
func AlitripMerchantGalaxyWechatCardQueryRecord(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatCardQueryRecordAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatCardQueryRecordAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
