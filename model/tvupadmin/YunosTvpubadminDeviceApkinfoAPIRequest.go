package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceApkinfoAPIRequest
获取停开服apk信息 API请求
yunos.tvpubadmin.device.apkinfo

获取停开服apk信息 */
type YunosTvpubadminDeviceApkinfoAPIRequest struct {
	model.Params
	// apkid
	_id int64
}

// New
