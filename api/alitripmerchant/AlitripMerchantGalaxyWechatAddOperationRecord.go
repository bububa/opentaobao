package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatAddOperationRecord 用户领取会员卡记录接口
// alitrip.merchant.galaxy.wechat.add.operation.record
//
// 用户领取会员卡记录接口
func AlitripMerchantGalaxyWechatAddOperationRecord(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatAddOperationRecordAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
