package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateDeviceserviceSearchmodelsAPIRequest
根据关键词检索设备型号 API请求
yunos.osupdate.deviceservice.searchmodels

根据关键词检索设备型号 */
type YunosOsupdateDeviceserviceSearchmodelsAPIRequest struct {
	model.Params
	// 设备父ID
	_parentId int64
	// 关键词
	_name string
}

// New
