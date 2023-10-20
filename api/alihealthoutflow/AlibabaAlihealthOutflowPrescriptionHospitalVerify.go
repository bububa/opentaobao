package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthoutflowprescriptionhospitalverify 处方同步至医院返回校验结果
// alibaba.alihealth.outflow.prescription.hospital.verify
//
// 处方同步至医院返回校验结果
func Alibabaalihealthoutflowprescriptionhospitalverify(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthoutflowprescriptionhospitalverifyAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthoutflowprescriptionhospitalverifyAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthoutflowprescriptionhospitalverifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
