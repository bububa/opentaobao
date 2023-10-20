package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatCardParmQuery 微信会员卡添加
// alitrip.merchant.galaxy.wechat.card.parm.query
//
// 微信会员卡添加参数获取
func AlitripMerchantGalaxyWechatCardParmQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatCardParmQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatCardParmQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
