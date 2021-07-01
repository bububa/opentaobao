package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketReverseAPIRequest
电子凭证冲正接口 API请求
taobao.vmarket.eticket.reverse

电子凭证平台冲正接口 */
type TaobaoVmarketEticketReverseAPIRequest struct {
	model.Params
	// 进行验码的电子凭证订单的订单ID
	_orderId int64
	// 冲正的码，只支持单个码
	_reverseCode string
	// 冲正份数（必须是和被冲正的核销记录的份数一致）
	_reverseNum int64
	// 需要冲正的核销记录对应核销流水号（对应的核销操作时候传递的自定义流水号）
	_consumeSecialNum string
	// 所有冲正后需要重新生成的码和对应的次数。码和次数之间用英文冒号分隔，多个码之间用英文逗号分隔。如果冲正后不需要重新生成码，留空
	_verifyCodes string
	// 不需要上传二维码图片或者冲正后不需要变更码的请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
	_qrImages string
	// 安全验证token，需要和该订单发码通知中的token一致
	_token string
	// 码商ID，是码商的话必须传递，如果是信任卖家不要传
	_codemerchantId int64
	// 机具id，如果是码商必须传，如果是信任卖家不要传
	_posid string
}

// New
