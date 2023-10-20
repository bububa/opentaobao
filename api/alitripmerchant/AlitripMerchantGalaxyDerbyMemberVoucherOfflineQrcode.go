package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervoucherofflineqrcode 德比线下权益券二维码查询
// alitrip.merchant.galaxy.derby.member.voucher.offline.qrcode
//
// 德比线下权益券二维码查询
func Alitripmerchantgalaxyderbymembervoucherofflineqrcode(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
