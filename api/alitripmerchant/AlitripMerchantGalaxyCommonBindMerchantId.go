package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxycommonbindmerchantid 绑定用户和merchantID
// alitrip.merchant.galaxy.common.bind.merchant.id
//
// 绑定用户和merchantID
func Alitripmerchantgalaxycommonbindmerchantid(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxycommonbindmerchantidAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxycommonbindmerchantidAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxycommonbindmerchantidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
