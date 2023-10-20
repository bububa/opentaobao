package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcode 德比线下权益券二维码查询
// alitrip.merchant.galaxy.derby.member.voucher.offline.qrcode
//
// 德比线下权益券二维码查询
func AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcode(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
