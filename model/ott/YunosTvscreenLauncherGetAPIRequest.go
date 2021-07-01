package ott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvscreenLauncherGetAPIRequest
一体机桌面 API请求
yunos.tvscreen.launcher.get

LCTS一体机桌面后台,提供基于运营坑位适配的桌面服务 */
type YunosTvscreenLauncherGetAPIRequest struct {
	model.Params
	// 设备属性
	_property string
	// IP来源
	_ip string
}

// New
