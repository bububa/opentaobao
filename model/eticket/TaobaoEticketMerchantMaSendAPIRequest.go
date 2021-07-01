package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoEticketMerchantMaSendAPIRequest
码商发码成功回调接口 API请求
taobao.eticket.merchant.ma.send

码商发码成功回调接口 */
type TaobaoEticketMerchantMaSendAPIRequest struct {
	model.Params
	// 业务类型
	_bizType int64
	// 需要发送的码列表
	_isvMaList []IsvMa
	// 业务id（订单号）
	_outerId string
	// 需要跟发码通知获取到的参数一致
	_token string
}

// New
