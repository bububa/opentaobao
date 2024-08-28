package alihealthoutflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthDoctorIncomeUpdate 医蝶谷医生收入打款情况回调
// alibaba.alihealth.doctor.income.update
//
// 医蝶谷医生收入打款情况回调
func AlibabaAlihealthDoctorIncomeUpdate(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthDoctorIncomeUpdateAPIRequest, resp *alihealthoutflow.AlibabaAlihealthDoctorIncomeUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
