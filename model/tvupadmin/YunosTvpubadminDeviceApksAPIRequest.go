package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceApksAPIRequest
获取停开服apk列表 API请求
yunos.tvpubadmin.device.apks

获取停开服apk列表 */
type YunosTvpubadminDeviceApksAPIRequest struct {
	model.Params
	// 牌照
	_license int64
}

// New
