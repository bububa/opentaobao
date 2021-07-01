package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketSendAPIRequest
商家电子凭证发码成功回调接口 API请求
taobao.vmarket.eticket.send

外部商家成功发码回调接口 */
type TaobaoVmarketEticketSendAPIRequest struct {
	model.Params
	// 订单编号
	_orderId int64
	// 发送成功的验证码及可验证次数的列表，码和可验证次数用英文冒号分隔，多个码之间用英文逗号分隔，所有字符都为英文半角
	_verifyCodes string
	// 安全验证token，需要和发码通知中的token一致
	_token string
	// 码商ID,是码商的话必须传递,如果是信任卖家,不需要传
	_codemerchantId int64
	// 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
	_qrImages string
}

// New
