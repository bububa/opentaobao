package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxycitylike 星河-酒店城市模糊查询
// alitrip.merchant.galaxy.city.like
//
// 根据城市模糊查询，雅高酒店所在城市的城市信息
func Alitripmerchantgalaxycitylike(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxycitylikeAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxycitylikeAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxycitylikeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
