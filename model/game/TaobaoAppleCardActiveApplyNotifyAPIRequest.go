package game

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAppleCardActiveApplyNotifyAPIRequest
苹果卡密申请激活回调接口 API请求
taobao.apple.card.active.apply.notify

苹果卡密申请激活回调接口 */
type TaobaoAppleCardActiveApplyNotifyAPIRequest struct {
	model.Params
	// 卡列表
	_appleCards []AppleCardDto
	// 网关订单号
	_gatewayOrderNo string
	// 描述
	_resultMsg string
	// 商户唯一订单号
	_orderNo string
	// 结果，000：成功，其他皆为错误 04： 订单处理失败(商户可退款，其他不可退款)
	_resultCode string
}

// New
