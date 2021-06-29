package ioti

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子价签ota接口 APIRequest
alibaba.it.esl.sendota

厂测接口，电子价签ota接口
*/
type AlibabaItEslSendotaRequest struct {
    model.Params

    // mac
    macAp   string 

    // base64的ota包
    otaDataBase64String   string 

}

func NewAlibabaItEslSendotaRequest() *AlibabaItEslSendotaRequest{
    return &AlibabaItEslSendotaRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItEslSendotaRequest) GetApiMethodName() string {
    return "alibaba.it.esl.sendota"
}

func (r AlibabaItEslSendotaRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItEslSendotaRequest) SetMacAp(macAp string) error {
    r.macAp = macAp
    r.Set("mac_ap", macAp)
    return nil
}

func (r AlibabaItEslSendotaRequest) GetMacAp() string {
    return r.macAp
}

func (r *AlibabaItEslSendotaRequest) SetOtaDataBase64String(otaDataBase64String string) error {
    r.otaDataBase64String = otaDataBase64String
    r.Set("ota_data_base64_string", otaDataBase64String)
    return nil
}

func (r AlibabaItEslSendotaRequest) GetOtaDataBase64String() string {
    return r.otaDataBase64String
}

