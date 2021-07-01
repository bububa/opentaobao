package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketConsumeAPIRequest
电子票券消费通知 API请求
taobao.vmarket.eticket.consume

外部合作商家电子票券消费回调接口 */
type TaobaoVmarketEticketConsumeAPIRequest struct {
	model.Params
	// 进行验码的电子凭证订单的订单ID
	_orderId int64
	// 核销的码，只支持单个码，多个码核销需要多次调用
	_verifyCode string
	// 核销份数
	_consumeNum int64
	// 安全验证token,需要和发码通知中的token一致
	_token string
	// 码商ID,是码商的话必须传递,如果是信任卖家不需要传
	_codemerchantId int64
	// 机具ID(此参数信任卖家可不传递，码商必须传递)
	_posid string
	// 手机后四位(没有特殊说明请不要传该参数)
	_mobile string
	// 核销后需要重新生成的码，如果不需要重新生成码，不要传该参数
	_newCode string
	// 自定义核销流水号，如果核销调用失败，可以用该核销流水号进行冲正操作，需要小于等于100个字符(a-zA-Z0-9_)；每次核销都是唯一的流水号
	_serialNum string
	// 不需要上传二维码图片或者核销后不需重新生成码码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），文件名与参数new_code必须相互对应。
	_qrImages string
}

// New
