package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDegoperationGetByEventkeyAPIRequest
通用用户抽奖次数限制 API请求
taobao.degoperation.get.by.eventkey

通用用户抽奖次数限制 */
type TaobaoDegoperationGetByEventkeyAPIRequest struct {
	model.Params
	// 活动后台配置appkey
	_degAppKey string
	// 活动后台配置eventkey
	_eventKey string
	// info
	_degAccessToken string
}

// New
