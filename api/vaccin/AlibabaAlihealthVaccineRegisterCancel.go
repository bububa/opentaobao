package vaccin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineRegisterCancel 取消登记
// alibaba.alihealth.vaccine.register.cancel
//
// 取消登记
func AlibabaAlihealthVaccineRegisterCancel(ctx context.Context, clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineRegisterCancelAPIRequest, resp *vaccin.AlibabaAlihealthVaccineRegisterCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
