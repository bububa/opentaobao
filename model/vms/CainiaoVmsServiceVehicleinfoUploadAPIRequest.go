package vms

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoVmsServiceVehicleinfoUploadAPIRequest 新能源车--外部车辆信息回传 API请求
// cainiao.vms.service.vehicleinfo.upload
//
// 新能源车--外部车辆信息回传
type CainiaoVmsServiceVehicleinfoUploadAPIRequest struct {
	model.Params
	// 设备号
	_deviceId string
	// 厂家标识
	_providerName string
	// 数据源标识
	_dataSource string
	// 协议版本标识
	_protocolVersion string
	// 上传的信息
	_data string
}

// NewCainiaoVmsServiceVehicleinfoUploadRequest 初始化CainiaoVmsServiceVehicleinfoUploadAPIRequest对象
func NewCainiaoVmsServiceVehicleinfoUploadRequest() *CainiaoVmsServiceVehicleinfoUploadAPIRequest {
	return &CainiaoVmsServiceVehicleinfoUploadAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoVmsServiceVehicleinfoUploadAPIRequest) Reset() {
	r._deviceId = ""
	r._providerName = ""
	r._dataSource = ""
	r._protocolVersion = ""
	r._data = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoVmsServiceVehicleinfoUploadAPIRequest) GetApiMethodName() string {
	return "cainiao.vms.service.vehicleinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoVmsServiceVehicleinfoUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoVmsServiceVehicleinfoUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备号
func (r *CainiaoVmsServiceVehicleinfoUploadAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r CainiaoVmsServiceVehicleinfoUploadAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetProviderName is ProviderName Setter
// 厂家标识
func (r *CainiaoVmsServiceVehicleinfoUploadAPIRequest) SetProviderName(_providerName string) error {
	r._providerName = _providerName
	r.Set("provider_name", _providerName)
	return nil
}

// GetProviderName ProviderName Getter
func (r CainiaoVmsServiceVehicleinfoUploadAPIRequest) GetProviderName() string {
	return r._providerName
}

// SetDataSource is DataSource Setter
// 数据源标识
func (r *CainiaoVmsServiceVehicleinfoUploadAPIRequest) SetDataSource(_dataSource string) error {
	r._dataSource = _dataSource
	r.Set("data_source", _dataSource)
	return nil
}

// GetDataSource DataSource Getter
func (r CainiaoVmsServiceVehicleinfoUploadAPIRequest) GetDataSource() string {
	return r._dataSource
}

// SetProtocolVersion is ProtocolVersion Setter
// 协议版本标识
func (r *CainiaoVmsServiceVehicleinfoUploadAPIRequest) SetProtocolVersion(_protocolVersion string) error {
	r._protocolVersion = _protocolVersion
	r.Set("protocol_version", _protocolVersion)
	return nil
}

// GetProtocolVersion ProtocolVersion Getter
func (r CainiaoVmsServiceVehicleinfoUploadAPIRequest) GetProtocolVersion() string {
	return r._protocolVersion
}

// SetData is Data Setter
// 上传的信息
func (r *CainiaoVmsServiceVehicleinfoUploadAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r CainiaoVmsServiceVehicleinfoUploadAPIRequest) GetData() string {
	return r._data
}

var poolCainiaoVmsServiceVehicleinfoUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoVmsServiceVehicleinfoUploadRequest()
	},
}

// GetCainiaoVmsServiceVehicleinfoUploadRequest 从 sync.Pool 获取 CainiaoVmsServiceVehicleinfoUploadAPIRequest
func GetCainiaoVmsServiceVehicleinfoUploadAPIRequest() *CainiaoVmsServiceVehicleinfoUploadAPIRequest {
	return poolCainiaoVmsServiceVehicleinfoUploadAPIRequest.Get().(*CainiaoVmsServiceVehicleinfoUploadAPIRequest)
}

// ReleaseCainiaoVmsServiceVehicleinfoUploadAPIRequest 将 CainiaoVmsServiceVehicleinfoUploadAPIRequest 放入 sync.Pool
func ReleaseCainiaoVmsServiceVehicleinfoUploadAPIRequest(v *CainiaoVmsServiceVehicleinfoUploadAPIRequest) {
	v.Reset()
	poolCainiaoVmsServiceVehicleinfoUploadAPIRequest.Put(v)
}
