package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest
设备状态查询 API请求
alibaba.ailabs.tmallgenie.auth.device.status.get

提供给天猫精灵定制机厂商 查询设备在线状态值 */
type AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest struct {
	model.Params
	// 给产品分配的唯一ID（22位）
	_clientId string
	// 精灵用户的唯一外部ID
	_userOpenId string
	// 精灵设备的唯一ID
	_uuid string
}

// New
