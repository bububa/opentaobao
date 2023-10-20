package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityPopupQuery 星河-获取雅高小程序营销抽奖首页弹窗
// alitrip.merchant.galaxy.activity.popup.query
//
// 获取雅高微信小程序，营销抽奖首页弹窗数据。
func AlitripMerchantGalaxyActivityPopupQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityPopupQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyActivityPopupQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
