package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaDeviceIsonlineAPIRequest
根据设备id查询设备是否在线 API请求
yunos.service.cmns.coa.device.isonline

根据设备id查询设备是否在线 */
type YunosServiceCmnsCoaDeviceIsonlineAPIRequest struct {
	model.Params
	// 设备id类型，取值"uuid"或者"imei"或者"deviceToken"
	_type string
	// 对应的设备id值
	_value string
}

// New
