package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallback 德比无限次券核销通知接口
// alitrip.merchant.galaxy.derby.voucher.card.unlimited.change.callback
//
// 德比无限次券核销通知接口
func AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallback(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
