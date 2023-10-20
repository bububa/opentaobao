package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlPlaybyid 通过id播放歌曲
// taobao.ailab.aicloud.top.device.control.playbyid
//
// 通过id播放歌曲
func TaobaoAilabAicloudTopDeviceControlPlaybyid(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
