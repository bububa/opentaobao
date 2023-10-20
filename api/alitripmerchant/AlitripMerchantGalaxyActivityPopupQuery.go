package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityPopupQuery 星河-获取雅高小程序营销抽奖首页弹窗
// alitrip.merchant.galaxy.activity.popup.query
//
// 获取雅高微信小程序，营销抽奖首页弹窗数据。
func AlitripMerchantGalaxyActivityPopupQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityPopupQueryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyActivityPopupQueryAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyActivityPopupQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
