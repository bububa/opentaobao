package smartstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

// TaobaoSmartstoreDeviceShelfFeedback 智能硬件云货架数据回流
// taobao.smartstore.device.shelf.feedback
//
// 智慧门店云货架设备回流
// 规则：
// 1. 回流的设备属于当前授权的用户
// 2. 回流的设备属于当前应用添加
func TaobaoSmartstoreDeviceShelfFeedback(clt *core.SDKClient, req *smartstore.TaobaoSmartstoreDeviceShelfFeedbackAPIRequest, session string) (*smartstore.TaobaoSmartstoreDeviceShelfFeedbackAPIResponse, error) {
	var resp smartstore.TaobaoSmartstoreDeviceShelfFeedbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
