package vms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新能源车--外部车辆信息回传 API请求
cainiao.vms.service.vehicleinfo.upload

新能源车--外部车辆信息回传
*/
type CainiaoVmsServiceVehicleinfoUploadRequest struct {
    model.Params
    // 设备号
    _deviceId   string
    // 厂家标识
    _providerName   string
    // 数据源标识
    _dataSource   string
    // 协议版本标识
    _protocolVersion   string
    // 上传的信息
    _data   string
}

// 初始化CainiaoVmsServiceVehicleinfoUploadRequest对象
func NewCainiaoVmsServiceVehicleinfoUploadRequest() *CainiaoVmsServiceVehicleinfoUploadRequest{
    return &CainiaoVmsServiceVehicleinfoUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetApiMethodName() string {
    return "cainiao.vms.service.vehicleinfo.upload"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备号
func (r *CainiaoVmsServiceVehicleinfoUploadRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetDeviceId() string {
    return r._deviceId
}
// ProviderName Setter
// 厂家标识
func (r *CainiaoVmsServiceVehicleinfoUploadRequest) SetProviderName(_providerName string) error {
    r._providerName = _providerName
    r.Set("provider_name", _providerName)
    return nil
}

// ProviderName Getter
func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetProviderName() string {
    return r._providerName
}
// DataSource Setter
// 数据源标识
func (r *CainiaoVmsServiceVehicleinfoUploadRequest) SetDataSource(_dataSource string) error {
    r._dataSource = _dataSource
    r.Set("data_source", _dataSource)
    return nil
}

// DataSource Getter
func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetDataSource() string {
    return r._dataSource
}
// ProtocolVersion Setter
// 协议版本标识
func (r *CainiaoVmsServiceVehicleinfoUploadRequest) SetProtocolVersion(_protocolVersion string) error {
    r._protocolVersion = _protocolVersion
    r.Set("protocol_version", _protocolVersion)
    return nil
}

// ProtocolVersion Getter
func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetProtocolVersion() string {
    return r._protocolVersion
}
// Data Setter
// 上传的信息
func (r *CainiaoVmsServiceVehicleinfoUploadRequest) SetData(_data string) error {
    r._data = _data
    r.Set("data", _data)
    return nil
}

// Data Getter
func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetData() string {
    return r._data
}
