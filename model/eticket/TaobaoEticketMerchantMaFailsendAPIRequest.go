package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoEticketMerchantMaFailsendAPIRequest
码商发码失败回调接口 API请求
taobao.eticket.merchant.ma.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证 */
type TaobaoEticketMerchantMaFailsendAPIRequest struct {
	model.Params
	// 业务id（订单号）
	_outerId string
	// 错误原因码
	_subErrCode string
	// 错误码描述
	_subErrMsg string
	// 需要与发码通知获取的值一致
	_token string
	// 业务类型
	_bizType int64
}

// New
