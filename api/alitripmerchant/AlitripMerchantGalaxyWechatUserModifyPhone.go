package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxywechatusermodifyphone DFC-ID用户换绑手机号
// alitrip.merchant.galaxy.wechat.user.modify.phone
//
// DFC-ID用户换绑手机号
func Alitripmerchantgalaxywechatusermodifyphone(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxywechatusermodifyphoneAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxywechatusermodifyphoneAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxywechatusermodifyphoneAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
