package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCardInfo 获取会员体系
// alitrip.merchant.galaxy.card.info
//
// 星河=根据卡类型获取当前的会员体系
func AlitripMerchantGalaxyCardInfo(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCardInfoAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyCardInfoAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyCardInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
