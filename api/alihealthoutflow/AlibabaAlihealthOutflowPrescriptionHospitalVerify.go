package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowPrescriptionHospitalVerify 处方同步至医院返回校验结果
// alibaba.alihealth.outflow.prescription.hospital.verify
//
// 处方同步至医院返回校验结果
func AlibabaAlihealthOutflowPrescriptionHospitalVerify(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest, resp *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
