package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxymembercard 星河-获取会员卡信息
// alitrip.merchant.galaxy.member.card
//
// 星河=根据会员等级获取会员的权益
func Alitripmerchantgalaxymembercard(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxymembercardAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxymembercardAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxymembercardAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
