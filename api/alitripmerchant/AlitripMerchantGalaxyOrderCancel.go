package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderCancel 星河-取消预订
// alitrip.merchant.galaxy.order.cancel
//
// 雅高酒店用户使用该接口，取消酒店预订
func AlitripMerchantGalaxyOrderCancel(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderCancelAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyOrderCancelAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyOrderCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
