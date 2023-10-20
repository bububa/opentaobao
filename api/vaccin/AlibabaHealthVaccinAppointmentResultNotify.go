package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinAppointmentResultNotify 通知预约结果
// alibaba.health.vaccin.appointment.result.notify
//
// 和ISV合作，需ISV回传预约结果。
func AlibabaHealthVaccinAppointmentResultNotify(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinAppointmentResultNotifyAPIRequest, resp *vaccin.AlibabaHealthVaccinAppointmentResultNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
