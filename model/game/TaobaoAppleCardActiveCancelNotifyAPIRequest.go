package game

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAppleCardActiveCancelNotifyAPIRequest
苹果卡密取消激活回调接口 API请求
taobao.apple.card.active.cancel.notify

苹果卡密取消激活回调接口 */
type TaobaoAppleCardActiveCancelNotifyAPIRequest struct {
	model.Params
	// 原商户申请激活时的订单号
	_orderNo string
	// 结果，000：整单取消激活成功  04： 取消激活失败（包括全部失败和部分失败，此时需以detail为准）
	_resultCode string
	// 网关订单号
	_gatewayOrderNo string
	// 描述
	_resultMsg string
	// 卡信息
	_appleCardCancels []AppleCardCancelDto
}

// New
