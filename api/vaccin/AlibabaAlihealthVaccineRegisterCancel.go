package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineRegisterCancel 取消登记
// alibaba.alihealth.vaccine.register.cancel
//
// 取消登记
func AlibabaAlihealthVaccineRegisterCancel(clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineRegisterCancelAPIRequest, resp *vaccin.AlibabaAlihealthVaccineRegisterCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
