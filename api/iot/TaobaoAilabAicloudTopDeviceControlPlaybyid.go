package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicecontrolplaybyid 通过id播放歌曲
// taobao.ailab.aicloud.top.device.control.playbyid
//
// 通过id播放歌曲
func Taobaoailabaicloudtopdevicecontrolplaybyid(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicecontrolplaybyidAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicecontrolplaybyidAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicecontrolplaybyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
