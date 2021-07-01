package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketResendAPIRequest
外部合作商家重发电子凭证回调接口 API请求
taobao.vmarket.eticket.resend

外部合作商家重发电子凭证回调接口 */
type TaobaoVmarketEticketResendAPIRequest struct {
	model.Params
	// 订单编号
	_orderId int64
	// 重新发送的验证码及可验证次数的列表，多个码之间用英文逗号分割，需要包含此订单所有可用的码（如果订单总的有10个码，可用的是5个，那么这里设置的是5个可用的码）
	_verifyCodes string
	// 安全验证token,回传淘宝发通知时发过来的token串
	_token string
	// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
	_codemerchantId int64
	// 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
	_qrImages string
}

// New
