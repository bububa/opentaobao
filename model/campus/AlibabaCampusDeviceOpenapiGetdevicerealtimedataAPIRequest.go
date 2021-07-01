package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest
获取指定设备下指定参数的实时值 API请求
alibaba.campus.device.openapi.getdevicerealtimedata

获取指定设备下指定参数的实时值 */
type AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest struct {
	model.Params
	// 请求端信息
	_workBenchContext *WorkBenchContext
	// 设备uuid
	_uuid string
	// 参数code,如灯亮度参数为brightness;参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’[根据设备类型查看该设备所拥有的采集类参数]。
	_propertyCode string
}

// New
