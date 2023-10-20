package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdeviceopenapifeedbackeventinfo IVS事件处理反馈接口
// alibaba.campus.device.openapi.feedbackeventinfo
//
// 提供给第三方ISV的的事件信息处理反馈的接口
func Alibabacampusdeviceopenapifeedbackeventinfo(clt *core.SDKClient, req *campus.AlibabacampusdeviceopenapifeedbackeventinfoAPIRequest, session string) (*campus.AlibabacampusdeviceopenapifeedbackeventinfoAPIResponse, error) {
	var resp campus.AlibabacampusdeviceopenapifeedbackeventinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
