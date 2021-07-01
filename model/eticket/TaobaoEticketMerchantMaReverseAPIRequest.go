package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoEticketMerchantMaReverseAPIRequest
电子凭证冲正接口 API请求
taobao.eticket.merchant.ma.reverse

电子凭证平台冲正接口 */
type TaobaoEticketMerchantMaReverseAPIRequest struct {
	model.Params
	// 业务类型
	_bizType int64
	// 码值
	_code string
	// 业务id（订单号）
	_outerId string
	// 机具编号，如果核销时有则必传
	_posId string
	// 冲正份数，需要与核销份数一致
	_reverseNum int64
	// 需要冲正的核销序列号
	_serialNum string
	// 需要跟发码通知获取到的参数一致
	_token string
}

// New
