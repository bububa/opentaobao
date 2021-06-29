package alihealthoutflow

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthoutflow"
)

/* 
处方同步至医院返回校验结果 
alibaba.alihealth.outflow.prescription.hospital.verify

处方同步至医院返回校验结果
*/
func AlibabaAlihealthOutflowPrescriptionHospitalVerify(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest, session string) (*alihealthoutflow.AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse, error) {
    var resp alihealthoutflow.AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
