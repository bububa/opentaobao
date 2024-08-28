package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiFeedbackeventinfo IVS事件处理反馈接口
// alibaba.campus.device.openapi.feedbackeventinfo
//
// 提供给第三方ISV的的事件信息处理反馈的接口
func AlibabaCampusDeviceOpenapiFeedbackeventinfo(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
