package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyactivitypopupquery 星河-获取雅高小程序营销抽奖首页弹窗
// alitrip.merchant.galaxy.activity.popup.query
//
// 获取雅高微信小程序，营销抽奖首页弹窗数据。
func Alitripmerchantgalaxyactivitypopupquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyactivitypopupqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyactivitypopupqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyactivitypopupqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
