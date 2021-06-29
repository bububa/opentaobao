package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机开始核销接口 API请求
alibaba.retail.electronic.certificate.pre.confirm

零售终端贩卖机开始核销接口,返回待领的商品ID
*/
type AlibabaRetailElectronicCertificatePreConfirmRequest struct {
    model.Params
    // 设备ID
    deviceId   string
    // 核销码
    code   int64
}

// 初始化AlibabaRetailElectronicCertificatePreConfirmRequest对象
func NewAlibabaRetailElectronicCertificatePreConfirmRequest() *AlibabaRetailElectronicCertificatePreConfirmRequest{
    return &AlibabaRetailElectronicCertificatePreConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailElectronicCertificatePreConfirmRequest) GetApiMethodName() string {
    return "alibaba.retail.electronic.certificate.pre.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailElectronicCertificatePreConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备ID
func (r *AlibabaRetailElectronicCertificatePreConfirmRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaRetailElectronicCertificatePreConfirmRequest) GetDeviceId() string {
    return r.deviceId
}
// Code Setter
// 核销码
func (r *AlibabaRetailElectronicCertificatePreConfirmRequest) SetCode(code int64) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaRetailElectronicCertificatePreConfirmRequest) GetCode() int64 {
    return r.code
}
