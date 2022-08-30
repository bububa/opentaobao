package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthDoctorLeshuiApplyNotify 申请单审核结果通知
// alibaba.alihealth.doctor.leshui.apply.notify
//
// 申请单审核结果通知
func AlibabaAlihealthDoctorLeshuiApplyNotify(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthDoctorLeshuiApplyNotifyAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse, error) {
	var resp alihealthoutflow.AlibabaAlihealthDoctorLeshuiApplyNotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
