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
type AlibabaRetailElectronicCertificatePreConfirmAPIRequest struct {
    model.Params
    // 设备ID
    _deviceId   string
    // 核销码
    _code   int64
}

// 初始化AlibabaRetailElectronicCertificatePreConfirmAPIRequest对象
func NewAlibabaRetailElectronicCertificatePreConfirmRequest() *AlibabaRetailElectronicCertificatePreConfirmAPIRequest{
    return &AlibabaRetailElectronicCertificatePreConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailElectronicCertificatePreConfirmAPIRequest) GetApiMethodName() string {
    return "alibaba.retail.electronic.certificate.pre.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailElectronicCertificatePreConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备ID
func (r *AlibabaRetailElectronicCertificatePreConfirmAPIRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaRetailElectronicCertificatePreConfirmAPIRequest) GetDeviceId() string {
    return r._deviceId
}
// Code Setter
// 核销码
func (r *AlibabaRetailElectronicCertificatePreConfirmAPIRequest) SetCode(_code int64) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaRetailElectronicCertificatePreConfirmAPIRequest) GetCode() int64 {
    return r._code
}
