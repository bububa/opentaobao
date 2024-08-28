package alihealthoutflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthDoctorLeshuiAuditResult 乐税审核结果通知
// alibaba.alihealth.doctor.leshui.audit.result
//
// 乐税审核结果通知
func AlibabaAlihealthDoctorLeshuiAuditResult(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthDoctorLeshuiAuditResultAPIRequest, resp *alihealthoutflow.AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
