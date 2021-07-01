package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaDeviceGetAPIRequest
设备详情查询 API请求
yunos.service.cmns.coa.device.get

第三方应用开发者调用此接口查询设备详情 */
type YunosServiceCmnsCoaDeviceGetAPIRequest struct {
	model.Params
	// 设备id类型,可以是uuid,imei,deviceToken,kp
	_type string
	// 设备id
	_value string
}

// New
