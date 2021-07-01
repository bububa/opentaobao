package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoEticketMerchantMaResendAPIRequest
电子凭证重发回调接口 API请求
taobao.eticket.merchant.ma.resend

码商重发电子凭证回调接口 */
type TaobaoEticketMerchantMaResendAPIRequest struct {
	model.Params
	// 业务类型
	_bizType int64
	// 待重发的码列表
	_isvMaList []IsvMa
	// 业务id（订单号）
	_outerId string
	// 需要跟发码通知获取到的参数一致
	_token string
}

// New
