package alihealthcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthcrm"
)

/* 
互联网医院批量导入接口 
alibaba.alihealth.medicalbase.hospital.sync

互联网医院isv批量通过接口批量导入
*/
func AlibabaAlihealthMedicalbaseHospitalSync(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthMedicalbaseHospitalSyncRequest, session string) (*alihealthcrm.AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse, error) {
    var resp alihealthcrm.AlibabaAlihealthMedicalbaseHospitalSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
