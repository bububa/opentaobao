package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxycardinfo 获取会员体系
// alitrip.merchant.galaxy.card.info
//
// 星河=根据卡类型获取当前的会员体系
func Alitripmerchantgalaxycardinfo(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxycardinfoAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxycardinfoAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxycardinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
