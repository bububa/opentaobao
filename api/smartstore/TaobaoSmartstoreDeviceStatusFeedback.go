package smartstore

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

// TaobaoSmartstoreDeviceStatusFeedback 设备在线状态回流
// taobao.smartstore.device.status.feedback
//
// 智能硬件设备状态回流
func TaobaoSmartstoreDeviceStatusFeedback(ctx context.Context, clt *core.SDKClient, req *smartstore.TaobaoSmartstoreDeviceStatusFeedbackAPIRequest, resp *smartstore.TaobaoSmartstoreDeviceStatusFeedbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
