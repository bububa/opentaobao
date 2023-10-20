package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineRegisterSubmit cdc回传疫苗登记数据
// alibaba.alihealth.vaccine.register.submit
//
// cdc回传疫苗登记信息
func AlibabaAlihealthVaccineRegisterSubmit(clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineRegisterSubmitAPIRequest, resp *vaccin.AlibabaAlihealthVaccineRegisterSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
