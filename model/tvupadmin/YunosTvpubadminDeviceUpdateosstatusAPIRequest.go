package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceUpdateosstatusAPIRequest
更新系统版本审核状态 API请求
yunos.tvpubadmin.device.updateosstatus

更新系统版本审核状态 */
type YunosTvpubadminDeviceUpdateosstatusAPIRequest struct {
	model.Params
	// 升级ID
	_versionId int64
	// 牌照方
	_license int64
	// 审核状态
	_status string
	// 审核意见
	_auditComment string
}

// New
