package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthdoctorleshuiapplynotify 申请单审核结果通知
// alibaba.alihealth.doctor.leshui.apply.notify
//
// 申请单审核结果通知
func Alibabaalihealthdoctorleshuiapplynotify(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthdoctorleshuiapplynotifyAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthdoctorleshuiapplynotifyAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthdoctorleshuiapplynotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
