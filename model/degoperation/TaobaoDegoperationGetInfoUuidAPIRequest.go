package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDegoperationGetInfoUuidAPIRequest
根据uuid用户抽奖次数限制 API请求
taobao.degoperation.get.info.uuid

根据uuid用户抽奖次数限制 */
type TaobaoDegoperationGetInfoUuidAPIRequest struct {
	model.Params
	// 活动后台配置eventkey
	_degAppKey string
	// 活动后台配置appkey
	_degEventKey string
	// 设备id
	_uuid string
}

// New
