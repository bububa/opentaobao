package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinappointmentresultnotify 通知预约结果
// alibaba.health.vaccin.appointment.result.notify
//
// 和ISV合作，需ISV回传预约结果。
func Alibabahealthvaccinappointmentresultnotify(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinappointmentresultnotifyAPIRequest, session string) (*vaccin.AlibabahealthvaccinappointmentresultnotifyAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinappointmentresultnotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
