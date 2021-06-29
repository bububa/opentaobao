package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认核销接口 API请求
alibaba.retail.electronic.certificate.confirm

确认核销接口
*/
type AlibabaRetailElectronicCertificateConfirmRequest struct {
    model.Params
    // 核销码
    _code   int64
    // 商品ID
    _itemId   int64
    // 设备ID
    _deviceId   string
}

// 初始化AlibabaRetailElectronicCertificateConfirmRequest对象
func NewAlibabaRetailElectronicCertificateConfirmRequest() *AlibabaRetailElectronicCertificateConfirmRequest{
    return &AlibabaRetailElectronicCertificateConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailElectronicCertificateConfirmRequest) GetApiMethodName() string {
    return "alibaba.retail.electronic.certificate.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailElectronicCertificateConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 核销码
func (r *AlibabaRetailElectronicCertificateConfirmRequest) SetCode(_code int64) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaRetailElectronicCertificateConfirmRequest) GetCode() int64 {
    return r._code
}
// ItemId Setter
// 商品ID
func (r *AlibabaRetailElectronicCertificateConfirmRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlibabaRetailElectronicCertificateConfirmRequest) GetItemId() int64 {
    return r._itemId
}
// DeviceId Setter
// 设备ID
func (r *AlibabaRetailElectronicCertificateConfirmRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaRetailElectronicCertificateConfirmRequest) GetDeviceId() string {
    return r._deviceId
}
