package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDevicehubOpenapiReportdataAPIRequest
设备数据上报 API请求
alibaba.campus.devicehub.openapi.reportdata

设备数据上报 */
type AlibabaCampusDevicehubOpenapiReportdataAPIRequest struct {
	model.Params
	// 自动生成
	_deviceEventData *DeviceReportEventDto
}

// New
