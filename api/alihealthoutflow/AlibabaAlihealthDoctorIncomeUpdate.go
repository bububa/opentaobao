package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthdoctorincomeupdate 医蝶谷医生收入打款情况回调
// alibaba.alihealth.doctor.income.update
//
// 医蝶谷医生收入打款情况回调
func Alibabaalihealthdoctorincomeupdate(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthdoctorincomeupdateAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthdoctorincomeupdateAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthdoctorincomeupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
