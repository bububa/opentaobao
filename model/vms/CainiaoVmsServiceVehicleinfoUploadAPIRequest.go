package vms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaovmsservicevehicleinfouploadAPIRequest 新能源车--外部车辆信息回传 API请求
// cainiao.vms.service.vehicleinfo.upload
//
// 新能源车--外部车辆信息回传
type CainiaovmsservicevehicleinfouploadAPIRequest struct {
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

// NewCainiaovmsservicevehicleinfouploadRequest 初始化CainiaovmsservicevehicleinfouploadAPIRequest对象
func NewCainiaovmsservicevehicleinfouploadRequest() *CainiaovmsservicevehicleinfouploadAPIRequest {
	return &CainiaovmsservicevehicleinfouploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaovmsservicevehicleinfouploadAPIRequest) GetApiMethodName() string {
	return "cainiao.vms.service.vehicleinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaovmsservicevehicleinfouploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaovmsservicevehicleinfouploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备号
func (r *CainiaovmsservicevehicleinfouploadAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r CainiaovmsservicevehicleinfouploadAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetProviderName is ProviderName Setter
// 厂家标识
func (r *CainiaovmsservicevehicleinfouploadAPIRequest) SetProviderName(_providerName string) error {
	r._providerName = _providerName
	r.Set("provider_name", _providerName)
	return nil
}

// GetProviderName ProviderName Getter
func (r CainiaovmsservicevehicleinfouploadAPIRequest) GetProviderName() string {
	return r._providerName
}

// SetDataSource is DataSource Setter
// 数据源标识
func (r *CainiaovmsservicevehicleinfouploadAPIRequest) SetDataSource(_dataSource string) error {
	r._dataSource = _dataSource
	r.Set("data_source", _dataSource)
	return nil
}

// GetDataSource DataSource Getter
func (r CainiaovmsservicevehicleinfouploadAPIRequest) GetDataSource() string {
	return r._dataSource
}

// SetProtocolVersion is ProtocolVersion Setter
// 协议版本标识
func (r *CainiaovmsservicevehicleinfouploadAPIRequest) SetProtocolVersion(_protocolVersion string) error {
	r._protocolVersion = _protocolVersion
	r.Set("protocol_version", _protocolVersion)
	return nil
}

// GetProtocolVersion ProtocolVersion Getter
func (r CainiaovmsservicevehicleinfouploadAPIRequest) GetProtocolVersion() string {
	return r._protocolVersion
}

// SetData is Data Setter
// 上传的信息
func (r *CainiaovmsservicevehicleinfouploadAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r CainiaovmsservicevehicleinfouploadAPIRequest) GetData() string {
	return r._data
}
