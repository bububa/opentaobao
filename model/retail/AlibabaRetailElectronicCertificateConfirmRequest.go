package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认核销接口 APIRequest
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

func NewAlibabaRetailElectronicCertificateConfirmRequest() *AlibabaRetailElectronicCertificateConfirmRequest{
    return &AlibabaRetailElectronicCertificateConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailElectronicCertificateConfirmRequest) GetApiMethodName() string {
    return "alibaba.retail.electronic.certificate.confirm"
}

func (r AlibabaRetailElectronicCertificateConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailElectronicCertificateConfirmRequest) SetCode(code int64) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaRetailElectronicCertificateConfirmRequest) GetCode() int64 {
    return r.code
}

func (r *AlibabaRetailElectronicCertificateConfirmRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaRetailElectronicCertificateConfirmRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlibabaRetailElectronicCertificateConfirmRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r AlibabaRetailElectronicCertificateConfirmRequest) GetDeviceId() string {
    return r.deviceId
}

