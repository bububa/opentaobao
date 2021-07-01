package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketBeforeconsumeAPIRequest
电子凭证验码前置确认 API请求
taobao.vmarket.eticket.beforeconsume

商家验码之前的调用接口，用来同步到最新的订单状态并判断是否可以进行验码操作 */
type TaobaoVmarketEticketBeforeconsumeAPIRequest struct {
	model.Params
	// 需要验码的电子凭证订单ID
	_orderId int64
	// 需要验的码
	_verifyCode string
	// 安全验证token，需要和发码通知中的token一致
	_token string
	// 码商ID,是码商的话必须传递,如果是信任卖家不需要传
	_codemerchantId int64
	// 操作员身份ID，如果是码商必须传,如果是信任卖家不需要传
	_posid string
	// 手机号码后四位,没有特殊说明请不要传
	_mobile string
}

// New
