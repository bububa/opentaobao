package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthDoctorLeshuiAuditResult 乐税审核结果通知
// alibaba.alihealth.doctor.leshui.audit.result
//
// 乐税审核结果通知
func AlibabaAlihealthDoctorLeshuiAuditResult(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse, error) {
	var resp alihealthoutflow.AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
