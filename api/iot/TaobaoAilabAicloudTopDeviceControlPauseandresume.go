package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicecontrolpauseandresume 设备播放暂停
// taobao.ailab.aicloud.top.device.control.pauseandresume
//
// 设备播放暂停
func Taobaoailabaicloudtopdevicecontrolpauseandresume(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicecontrolpauseandresumeAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicecontrolpauseandresumeAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicecontrolpauseandresumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
