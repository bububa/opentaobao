package vms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新能源车--外部车辆信息回传 APIRequest
cainiao.vms.service.vehicleinfo.upload

新能源车--外部车辆信息回传
*/
type CainiaoVmsServiceVehicleinfoUploadRequest struct {
    model.Params

    // 设备号
    deviceId   string 

    // 厂家标识
    providerName   string 

    // 数据源标识
    dataSource   string 

    // 协议版本标识
    protocolVersion   string 

    // 上传的信息
    data   string 

}

func NewCainiaoVmsServiceVehicleinfoUploadRequest() *CainiaoVmsServiceVehicleinfoUploadRequest{
    return &CainiaoVmsServiceVehicleinfoUploadRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetApiMethodName() string {
    return "cainiao.vms.service.vehicleinfo.upload"
}

func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoVmsServiceVehicleinfoUploadRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *CainiaoVmsServiceVehicleinfoUploadRequest) SetProviderName(providerName string) error {
    r.providerName = providerName
    r.Set("provider_name", providerName)
    return nil
}

func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetProviderName() string {
    return r.providerName
}

func (r *CainiaoVmsServiceVehicleinfoUploadRequest) SetDataSource(dataSource string) error {
    r.dataSource = dataSource
    r.Set("data_source", dataSource)
    return nil
}

func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetDataSource() string {
    return r.dataSource
}

func (r *CainiaoVmsServiceVehicleinfoUploadRequest) SetProtocolVersion(protocolVersion string) error {
    r.protocolVersion = protocolVersion
    r.Set("protocol_version", protocolVersion)
    return nil
}

func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetProtocolVersion() string {
    return r.protocolVersion
}

func (r *CainiaoVmsServiceVehicleinfoUploadRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r CainiaoVmsServiceVehicleinfoUploadRequest) GetData() string {
    return r.data
}

