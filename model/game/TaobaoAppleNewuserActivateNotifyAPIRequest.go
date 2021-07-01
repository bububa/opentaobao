package game

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAppleNewuserActivateNotifyAPIRequest
新用户激活通知接口 API请求
taobao.apple.newuser.activate.notify

资和信主动通知激活结果 */
type TaobaoAppleNewuserActivateNotifyAPIRequest struct {
	model.Params
	// 结果对应值，00位成功，其他为失败
	_resultCode string
	// 处理结果中文描述
	_resultMsg string
	// 主业务参数
	_mainData *AppleTopActivateNotifyDo
}

// New
