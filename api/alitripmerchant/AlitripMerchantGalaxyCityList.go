package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxycitylist 星河-酒店城市列表展示
// alitrip.merchant.galaxy.city.list
//
// 雅高酒店城市列表展示，并且首字母列出酒店城市
func Alitripmerchantgalaxycitylist(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxycitylistAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxycitylistAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxycitylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
