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
    code   int64
    // 商品ID
    itemId   int64
    // 设备ID
    deviceId   string
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
func (r *AlibabaRetailElectronicCertificateConfirmRequest) SetCode(code int64) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaRetailElectronicCertificateConfirmRequest) GetCode() int64 {
    return r.code
}
// ItemId Setter
// 商品ID
func (r *AlibabaRetailElectronicCertificateConfirmRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaRetailElectronicCertificateConfirmRequest) GetItemId() int64 {
    return r.itemId
}
// DeviceId Setter
// 设备ID
func (r *AlibabaRetailElectronicCertificateConfirmRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaRetailElectronicCertificateConfirmRequest) GetDeviceId() string {
    return r.deviceId
}
