package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatAddOperationRecord 用户领取会员卡记录接口
// alitrip.merchant.galaxy.wechat.add.operation.record
//
// 用户领取会员卡记录接口
func AlitripMerchantGalaxyWechatAddOperationRecord(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatAddOperationRecordAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatAddOperationRecordAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
