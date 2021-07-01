package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceModelsAPIRequest
获取品牌下设备列表 API请求
yunos.tvpubadmin.device.models

获取品牌下设备列表 */
type YunosTvpubadminDeviceModelsAPIRequest struct {
	model.Params
	// 终端类型
	_terminalType string
	// 品牌ID
	_brandId int64
	// 牌照方
	_license int64
}

// New
