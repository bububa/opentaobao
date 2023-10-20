package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDevicehubOpenapiReportdataAPIRequest 设备数据上报 API请求
// alibaba.campus.devicehub.openapi.reportdata
//
// 设备数据上报
type AlibabaCampusDevicehubOpenapiReportdataAPIRequest struct {
	model.Params
	// 设备数据上报信息
	_deviceEventData *DeviceReportEventDto
}

// NewAlibabaCampusDevicehubOpenapiReportdataRequest 初始化AlibabaCampusDevicehubOpenapiReportdataAPIRequest对象
func NewAlibabaCampusDevicehubOpenapiReportdataRequest() *AlibabaCampusDevicehubOpenapiReportdataAPIRequest {
	return &AlibabaCampusDevicehubOpenapiReportdataAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusDevicehubOpenapiReportdataAPIRequest) Reset() {
	r._deviceEventData = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDevicehubOpenapiReportdataAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.devicehub.openapi.reportdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDevicehubOpenapiReportdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusDevicehubOpenapiReportdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceEventData is DeviceEventData Setter
// 设备数据上报信息
func (r *AlibabaCampusDevicehubOpenapiReportdataAPIRequest) SetDeviceEventData(_deviceEventData *DeviceReportEventDto) error {
	r._deviceEventData = _deviceEventData
	r.Set("device_event_data", _deviceEventData)
	return nil
}

// GetDeviceEventData DeviceEventData Getter
func (r AlibabaCampusDevicehubOpenapiReportdataAPIRequest) GetDeviceEventData() *DeviceReportEventDto {
	return r._deviceEventData
}

var poolAlibabaCampusDevicehubOpenapiReportdataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusDevicehubOpenapiReportdataRequest()
	},
}

// GetAlibabaCampusDevicehubOpenapiReportdataRequest 从 sync.Pool 获取 AlibabaCampusDevicehubOpenapiReportdataAPIRequest
func GetAlibabaCampusDevicehubOpenapiReportdataAPIRequest() *AlibabaCampusDevicehubOpenapiReportdataAPIRequest {
	return poolAlibabaCampusDevicehubOpenapiReportdataAPIRequest.Get().(*AlibabaCampusDevicehubOpenapiReportdataAPIRequest)
}

// ReleaseAlibabaCampusDevicehubOpenapiReportdataAPIRequest 将 AlibabaCampusDevicehubOpenapiReportdataAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusDevicehubOpenapiReportdataAPIRequest(v *AlibabaCampusDevicehubOpenapiReportdataAPIRequest) {
	v.Reset()
	poolAlibabaCampusDevicehubOpenapiReportdataAPIRequest.Put(v)
}
