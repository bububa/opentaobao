package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcode 德比线下权益券二维码查询
// alitrip.merchant.galaxy.derby.member.voucher.offline.qrcode
//
// 德比线下权益券二维码查询
func AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcode(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
