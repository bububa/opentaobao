package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailDeviceInventorySyncAPIRequest
库存同步接口 API请求
alibaba.retail.device.inventory.sync

商库存同步接口 */
type AlibabaRetailDeviceInventorySyncAPIRequest struct {
	model.Params
	// 设备类型
	_deviceType string
	// 设备Id
	_deviceId string
	// 系统自动生成
	_inventoryDtos []InventorySyncDto
	// 系统自动生成
	_deviceOption *InventorySyncOption
}

// New
