package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthDoctorIncomeUpdate 医蝶谷医生收入打款情况回调
// alibaba.alihealth.doctor.income.update
//
// 医蝶谷医生收入打款情况回调
func AlibabaAlihealthDoctorIncomeUpdate(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthDoctorIncomeUpdateAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthDoctorIncomeUpdateAPIResponse, error) {
	var resp alihealthoutflow.AlibabaAlihealthDoctorIncomeUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
