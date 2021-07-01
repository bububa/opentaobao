package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest
获取五道口设备管理信息 API请求
taobao.wdk.equipment.deviceadmin.deviceinfo.get

通过仓编码获取五道口设备管理信息 */
type TaobaoWdkEquipmentDeviceadminDeviceinfoGetAPIRequest struct {
	model.Params
	// 仓编码
	_warehouseCode string
	// 设备类型
	_deviceType int64
}

// New
