package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthDoctorLeshuiTicketValid 乐税token验证
// alibaba.alihealth.doctor.leshui.ticket.valid
//
// 乐税token验证
func AlibabaAlihealthDoctorLeshuiTicketValid(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthDoctorLeshuiTicketValidAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse, error) {
	var resp alihealthoutflow.AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
