package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbyvouchercardunlimitedchangecallback 德比无限次券核销通知接口
// alitrip.merchant.galaxy.derby.voucher.card.unlimited.change.callback
//
// 德比无限次券核销通知接口
func Alitripmerchantgalaxyderbyvouchercardunlimitedchangecallback(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
