package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthDoctorLeshuiTicketValid 乐税token验证
// alibaba.alihealth.doctor.leshui.ticket.valid
//
// 乐税token验证
func AlibabaAlihealthDoctorLeshuiTicketValid(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest, resp *alihealthoutflow.AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
