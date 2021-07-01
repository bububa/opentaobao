package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceAppupgradedetailAPIRequest
获取应用升级详情 API请求
yunos.tvpubadmin.device.appupgradedetail

获取应用升级详情 */
type YunosTvpubadminDeviceAppupgradedetailAPIRequest struct {
	model.Params
	// 应用升级的ID
	_versionId int64
	// 牌照方
	_license int64
}

// New
