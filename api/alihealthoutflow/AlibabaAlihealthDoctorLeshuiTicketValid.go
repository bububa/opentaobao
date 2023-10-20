package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthdoctorleshuiticketvalid 乐税token验证
// alibaba.alihealth.doctor.leshui.ticket.valid
//
// 乐税token验证
func Alibabaalihealthdoctorleshuiticketvalid(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthdoctorleshuiticketvalidAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthdoctorleshuiticketvalidAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthdoctorleshuiticketvalidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
