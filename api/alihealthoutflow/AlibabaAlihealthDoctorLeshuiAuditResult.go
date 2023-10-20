package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthdoctorleshuiauditresult 乐税审核结果通知
// alibaba.alihealth.doctor.leshui.audit.result
//
// 乐税审核结果通知
func Alibabaalihealthdoctorleshuiauditresult(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthdoctorleshuiauditresultAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthdoctorleshuiauditresultAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthdoctorleshuiauditresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
