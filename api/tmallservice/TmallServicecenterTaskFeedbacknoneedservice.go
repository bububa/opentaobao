package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecentertaskfeedbacknoneedservice 服务商反馈无需安装工单接口
// tmall.servicecenter.task.feedbacknoneedservice
//
// 服务商反馈无需安装工单接口
func Tmallservicecentertaskfeedbacknoneedservice(clt *core.SDKClient, req *tmallservice.TmallservicecentertaskfeedbacknoneedserviceAPIRequest, session string) (*tmallservice.TmallservicecentertaskfeedbacknoneedserviceAPIResponse, error) {
	var resp tmallservice.TmallservicecentertaskfeedbacknoneedserviceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
