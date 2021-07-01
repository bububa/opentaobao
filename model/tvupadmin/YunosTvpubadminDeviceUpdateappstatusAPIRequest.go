package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceUpdateappstatusAPIRequest
更新应用版本审核状态 API请求
yunos.tvpubadmin.device.updateappstatus

更新应用版本审核状态 */
type YunosTvpubadminDeviceUpdateappstatusAPIRequest struct {
	model.Params
	// 应用ID
	_versionId int64
	// 牌照方
	_license int64
	// 审核状态
	_status string
	// 审核意见
	_auditComment string
}

// New
