package smartstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

// Taobaosmartstoredevicestatusfeedback 设备在线状态回流
// taobao.smartstore.device.status.feedback
//
// 智能硬件设备状态回流
func Taobaosmartstoredevicestatusfeedback(clt *core.SDKClient, req *smartstore.TaobaosmartstoredevicestatusfeedbackAPIRequest, session string) (*smartstore.TaobaosmartstoredevicestatusfeedbackAPIResponse, error) {
	var resp smartstore.TaobaosmartstoredevicestatusfeedbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
