package game

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAppleNewuserSignNotifyNewversionAPIRequest
新用户签约结果通知接口v2 API请求
taobao.apple.newuser.sign.notify.newversion

资和信主动通知签约结果 */
type TaobaoAppleNewuserSignNotifyNewversionAPIRequest struct {
	model.Params
	// 结果code
	_resultCode string
	// 结果信息说明
	_resultMsg string
	// 业务参数
	_mainData *AppleTopNewSignNotifyDo
}

// New
