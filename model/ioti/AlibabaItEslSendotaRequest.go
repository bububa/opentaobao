package ioti

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子价签ota接口 API请求
alibaba.it.esl.sendota

厂测接口，电子价签ota接口
*/
type AlibabaItEslSendotaRequest struct {
    model.Params
    // mac
    _macAp   string
    // base64的ota包
    _otaDataBase64String   string
}

// 初始化AlibabaItEslSendotaRequest对象
func NewAlibabaItEslSendotaRequest() *AlibabaItEslSendotaRequest{
    return &AlibabaItEslSendotaRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItEslSendotaRequest) GetApiMethodName() string {
    return "alibaba.it.esl.sendota"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItEslSendotaRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MacAp Setter
// mac
func (r *AlibabaItEslSendotaRequest) SetMacAp(_macAp string) error {
    r._macAp = _macAp
    r.Set("mac_ap", _macAp)
    return nil
}

// MacAp Getter
func (r AlibabaItEslSendotaRequest) GetMacAp() string {
    return r._macAp
}
// OtaDataBase64String Setter
// base64的ota包
func (r *AlibabaItEslSendotaRequest) SetOtaDataBase64String(_otaDataBase64String string) error {
    r._otaDataBase64String = _otaDataBase64String
    r.Set("ota_data_base64_string", _otaDataBase64String)
    return nil
}

// OtaDataBase64String Getter
func (r AlibabaItEslSendotaRequest) GetOtaDataBase64String() string {
    return r._otaDataBase64String
}
