package smartstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

// TaobaoSmartstoreDeviceStatusFeedback 设备在线状态回流
// taobao.smartstore.device.status.feedback
//
// 智能硬件设备状态回流
func TaobaoSmartstoreDeviceStatusFeedback(clt *core.SDKClient, req *smartstore.TaobaoSmartstoreDeviceStatusFeedbackAPIRequest, session string) (*smartstore.TaobaoSmartstoreDeviceStatusFeedbackAPIResponse, error) {
	var resp smartstore.TaobaoSmartstoreDeviceStatusFeedbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
