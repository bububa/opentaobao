package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatCardParmQuery 微信会员卡添加
// alitrip.merchant.galaxy.wechat.card.parm.query
//
// 微信会员卡添加参数获取
func AlitripMerchantGalaxyWechatCardParmQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatCardParmQueryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyWechatCardParmQueryAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyWechatCardParmQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
