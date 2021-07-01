package game

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAppleNewuserSignNotifyAPIRequest
新用户签约通知接口 API请求
taobao.apple.newuser.sign.notify

用户付款成功后，资和信主动通知签约结果 */
type TaobaoAppleNewuserSignNotifyAPIRequest struct {
	model.Params
	// 结果code
	_resultCode string
	// 结果信息说明
	_resultMsg string
	// 业务参数
	_mainData *AppleTopNewSignNotifyDo
}

// New
