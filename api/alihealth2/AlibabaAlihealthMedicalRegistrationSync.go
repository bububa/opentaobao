package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
阿里健康支付宝挂号记录回传接口 
alibaba.alihealth.medical.registration.sync

阿里健康支付宝挂号记录回传接口
*/
func AlibabaAlihealthMedicalRegistrationSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalRegistrationSyncAPIRequest, session string) (*alihealth2.AlibabaAlihealthMedicalRegistrationSyncAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthMedicalRegistrationSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
