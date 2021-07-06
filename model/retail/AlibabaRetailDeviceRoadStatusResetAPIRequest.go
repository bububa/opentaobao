package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailDeviceRoadStatusResetAPIRequest 贩卖机货道解锁 API请求
// alibaba.retail.device.road.status.reset
//
// 贩卖机货道解锁
type AlibabaRetailDeviceRoadStatusResetAPIRequest struct {
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

// NewAlibabaRetailDeviceRoadStatusResetRequest 初始化AlibabaRetailDeviceRoadStatusResetAPIRequest对象
func NewAlibabaRetailDeviceRoadStatusResetRequest() *AlibabaRetailDeviceRoadStatusResetAPIRequest {
	return &AlibabaRetailDeviceRoadStatusResetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailDeviceRoadStatusResetAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.device.road.status.reset"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailDeviceRoadStatusResetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRoadNoList is RoadNoList Setter
// 货道编码
func (r *AlibabaRetailDeviceRoadStatusResetAPIRequest) SetRoadNoList(_roadNoList []string) error {
	r._roadNoList = _roadNoList
	r.Set("road_no_list", _roadNoList)
	return nil
}

// GetRoadNoList RoadNoList Getter
func (r AlibabaRetailDeviceRoadStatusResetAPIRequest) GetRoadNoList() []string {
	return r._roadNoList
}

// SetDeviceUuid is DeviceUuid Setter
// 设备外部编码
func (r *AlibabaRetailDeviceRoadStatusResetAPIRequest) SetDeviceUuid(_deviceUuid string) error {
	r._deviceUuid = _deviceUuid
	r.Set("device_uuid", _deviceUuid)
	return nil
}

// GetDeviceUuid DeviceUuid Getter
func (r AlibabaRetailDeviceRoadStatusResetAPIRequest) GetDeviceUuid() string {
	return r._deviceUuid
}

// SetDeviceCode is DeviceCode Setter
// 阿里设备编码
func (r *AlibabaRetailDeviceRoadStatusResetAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r AlibabaRetailDeviceRoadStatusResetAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetDeviceSn is DeviceSn Setter
// 阿里设备物理编码
func (r *AlibabaRetailDeviceRoadStatusResetAPIRequest) SetDeviceSn(_deviceSn string) error {
	r._deviceSn = _deviceSn
	r.Set("device_sn", _deviceSn)
	return nil
}

// GetDeviceSn DeviceSn Getter
func (r AlibabaRetailDeviceRoadStatusResetAPIRequest) GetDeviceSn() string {
	return r._deviceSn
}
