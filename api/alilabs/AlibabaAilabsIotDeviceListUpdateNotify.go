package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabsiotdevicelistupdatenotify 设备列表更新通知
// alibaba.ailabs.iot.device.list.update.notify
//
// 用于人工智能实验室IoT合作厂商上报三方接入IoT设备列表更新时的通知
func Alibabaailabsiotdevicelistupdatenotify(clt *core.SDKClient, req *alilabs.AlibabaailabsiotdevicelistupdatenotifyAPIRequest, session string) (*alilabs.AlibabaailabsiotdevicelistupdatenotifyAPIResponse, error) {
	var resp alilabs.AlibabaailabsiotdevicelistupdatenotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
