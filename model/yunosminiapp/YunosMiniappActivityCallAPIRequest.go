package yunosminiapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosMiniappActivityCallAPIRequest
调用活动接口 API请求
yunos.miniapp.activity.call

用于小程序调用活动接口 */
type YunosMiniappActivityCallAPIRequest struct {
	model.Params
	// 请求选项
	_options *Options
	// 设备id
	_deviceId string
	// 活动id
	_activityId string
}

// New
