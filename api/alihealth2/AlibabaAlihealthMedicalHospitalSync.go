package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
阿里将康支付宝挂号数据医院回传接口 
alibaba.alihealth.medical.hospital.sync

阿里将康支付宝挂号数据医院回传接口
*/
func AlibabaAlihealthMedicalHospitalSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalHospitalSyncRequest, session string) (*alihealth2.AlibabaAlihealthMedicalHospitalSyncAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthMedicalHospitalSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
