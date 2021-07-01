package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceOsupgradedetailAPIRequest
获取系统升级详情 API请求
yunos.tvpubadmin.device.osupgradedetail

获取系统升级详情 */
type YunosTvpubadminDeviceOsupgradedetailAPIRequest struct {
	model.Params
	// 版本ID
	_versionId int64
	// 牌照方
	_license int64
}

// New
