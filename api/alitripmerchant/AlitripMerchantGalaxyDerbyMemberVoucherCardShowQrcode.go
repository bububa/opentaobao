package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardshowqrcode 会员权益卡身份识别二维码图片
// alitrip.merchant.galaxy.derby.member.voucher.card.show.qrcode
//
// 会员权益卡身份识别二维码图片
func Alitripmerchantgalaxyderbymembervouchercardshowqrcode(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardshowqrcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
