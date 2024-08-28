package alihealthoutflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthDoctorLeshuiApplyNotify 申请单审核结果通知
// alibaba.alihealth.doctor.leshui.apply.notify
//
// 申请单审核结果通知
func AlibabaAlihealthDoctorLeshuiApplyNotify(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest, resp *alihealthoutflow.AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
