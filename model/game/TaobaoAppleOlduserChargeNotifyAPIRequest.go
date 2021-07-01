package game

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAppleOlduserChargeNotifyAPIRequest
老用户激活并兑换通知接口 API请求
taobao.apple.olduser.charge.notify

老用户激活并兑换通知接口 */
type TaobaoAppleOlduserChargeNotifyAPIRequest struct {
	model.Params
	// 结果code
	_resultCode string
	// 结果信息说明
	_resultMsg string
	// 业务参数
	_mainData *AppleTopOldSignNotifyDo
}

// New
