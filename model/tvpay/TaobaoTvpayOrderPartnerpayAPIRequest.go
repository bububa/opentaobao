package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTvpayOrderPartnerpayAPIRequest
tv支付第三方支付订单 API请求
taobao.tvpay.order.partnerpay

tv支付第三方发起并支付订单（使用设备授权） */
type TaobaoTvpayOrderPartnerpayAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 订单信息
	_data string
	// 支付方式
	_payType string
	// 牌照方
	_license string
}

// New
