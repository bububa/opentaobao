package alihealthoutflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowPrescriptionHospitalVerify 处方同步至医院返回校验结果
// alibaba.alihealth.outflow.prescription.hospital.verify
//
// 处方同步至医院返回校验结果
func AlibabaAlihealthOutflowPrescriptionHospitalVerify(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest, resp *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
