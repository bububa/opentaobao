package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretaildeviceroadstatusresetAPIRequest 贩卖机货道解锁 API请求
// alibaba.retail.device.road.status.reset
//
// 贩卖机货道解锁
type AlibabaretaildeviceroadstatusresetAPIRequest struct {
	model.Params
	// 货道编码
	_roadNoList []string
	// 设备外部编码
	_deviceUuid string
	// 阿里设备编码
	_deviceCode string
	// 阿里设备物理编码
	_deviceSn string
}

// NewAlibabaretaildeviceroadstatusresetRequest 初始化AlibabaretaildeviceroadstatusresetAPIRequest对象
func NewAlibabaretaildeviceroadstatusresetRequest() *AlibabaretaildeviceroadstatusresetAPIRequest {
	return &AlibabaretaildeviceroadstatusresetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretaildeviceroadstatusresetAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.device.road.status.reset"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretaildeviceroadstatusresetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretaildeviceroadstatusresetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRoadNoList is RoadNoList Setter
// 货道编码
func (r *AlibabaretaildeviceroadstatusresetAPIRequest) SetRoadNoList(_roadNoList []string) error {
	r._roadNoList = _roadNoList
	r.Set("road_no_list", _roadNoList)
	return nil
}

// GetRoadNoList RoadNoList Getter
func (r AlibabaretaildeviceroadstatusresetAPIRequest) GetRoadNoList() []string {
	return r._roadNoList
}

// SetDeviceUuid is DeviceUuid Setter
// 设备外部编码
func (r *AlibabaretaildeviceroadstatusresetAPIRequest) SetDeviceUuid(_deviceUuid string) error {
	r._deviceUuid = _deviceUuid
	r.Set("device_uuid", _deviceUuid)
	return nil
}

// GetDeviceUuid DeviceUuid Getter
func (r AlibabaretaildeviceroadstatusresetAPIRequest) GetDeviceUuid() string {
	return r._deviceUuid
}

// SetDeviceCode is DeviceCode Setter
// 阿里设备编码
func (r *AlibabaretaildeviceroadstatusresetAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r AlibabaretaildeviceroadstatusresetAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetDeviceSn is DeviceSn Setter
// 阿里设备物理编码
func (r *AlibabaretaildeviceroadstatusresetAPIRequest) SetDeviceSn(_deviceSn string) error {
	r._deviceSn = _deviceSn
	r.Set("device_sn", _deviceSn)
	return nil
}

// GetDeviceSn DeviceSn Getter
func (r AlibabaretaildeviceroadstatusresetAPIRequest) GetDeviceSn() string {
	return r._deviceSn
}
