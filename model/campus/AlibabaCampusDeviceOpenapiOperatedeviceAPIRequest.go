package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest
根据uuid操作设备 API请求
alibaba.campus.device.openapi.operatedevice

根据uuid操作设备 */
type AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 设备序列号uuid
	_uuid string
	// 参数code,如灯亮度参数为brightness;设备的开关switchstate。参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’。
	_propertyCode string
	// 设置的参数值.如灯亮度为0~255.0表示关;设备开关,值使用on或off。[请按照‘设备详细信息开发文档’传入正确的参数值类型]
	_value string
}

// New
